/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"

	apps "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	websitev1alpha1 "openeuler.org/website-manager/api/v1alpha1"
)

// RepoScannerReconciler reconciles a RepoScanner object
type RepoScannerReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=website.openeuler.org,resources=reposcanners,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=website.openeuler.org,resources=reposcanners/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=website.openeuler.org,resources=reposcanners/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the RepoScanner object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.8.3/pkg/reconcile
func (r *RepoScannerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_log := log.FromContext(ctx)

	instance := websitev1alpha1.RepoScanner{}

	if err := r.Client.Get(ctx, req.NamespacedName, &instance); err != nil {
		//_log.Error(err, "failed to get RepoScanner resource")
		// Ignore NotFound errors as they will be retried automatically if the
		// resource is created in future.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	deployment := apps.Deployment{}
	err := r.Client.Get(ctx, client.ObjectKey{Namespace: instance.Namespace, Name: instance.Name + "-worker"}, &deployment)
	if apierrors.IsNotFound(err) {
		_log.Info("could not find existing Deployment for MyKind, creating one...")

		deployment = buildDeployment(instance)
		if err := r.Client.Create(ctx, &deployment); err != nil {

			_log.Error(err, "failed to create Deployment resource")
			return ctrl.Result{}, err
		}

		//r.Recorder.Eventf(&instance, v1.EventTypeNormal, "Created", "Created deployment %q", deployment.Name)
		_log.Info("created Deployment resource for MyKind")
		return ctrl.Result{}, nil
	}
	if err != nil {
		_log.Error(err, "failed to get Deployment for MyKind resource")
		return ctrl.Result{}, err
	}

	_log.Info("existing Deployment resource already exists for MyKind, checking replica count")

	expectedReplicas := int32(1)

	if *deployment.Spec.Replicas != expectedReplicas {
		_log.Info("updating replica count", "old_count", *deployment.Spec.Replicas, "new_count", expectedReplicas)

		deployment.Spec.Replicas = &expectedReplicas
		if err := r.Client.Update(ctx, &deployment); err != nil {
			_log.Error(err, "failed to Deployment update replica count")
			return ctrl.Result{}, err
		}

		//r.Recorder.Eventf(&instance, v1.EventTypeNormal, "Scaled", "Scaled deployment %q to %d replicas", deployment.Name, expectedReplicas)

		return ctrl.Result{}, nil
	}

	log.Log.Info("replica count up to date", "replica_count", *deployment.Spec.Replicas)

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *RepoScannerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&websitev1alpha1.RepoScanner{}).
		Owns(&apps.Deployment{}).
		Complete(r)
}

func buildDeployment(repoScanner websitev1alpha1.RepoScanner) apps.Deployment {

	var replicas int32 = 1

	deployment := apps.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:            repoScanner.Name + "-worker",
			Namespace:       repoScanner.Namespace,
			OwnerReferences: []metav1.OwnerReference{*metav1.NewControllerRef(&repoScanner, websitev1alpha1.GroupVersion.WithKind("RepoScanner"))},
		},
		Spec: apps.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"openeuler.org/website-operator-RepoScanner-Worker": repoScanner.Name + "-worker",
				},
			},
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"openeuler.org/website-operator-RepoScanner-Worker": repoScanner.Name + "-worker",
					},
				},
				Spec: repoScanner.Spec.PodSpec,
			},
		},
	}
	return deployment
}
