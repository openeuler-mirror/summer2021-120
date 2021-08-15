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

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// RepoScannerSpec defines the desired state of RepoScanner
type RepoScannerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of RepoScanner. Edit reposcanner_types.go to remove/update
	GiteeOwner       string         `json:"giteeOwner,omitempty"`
	GiteeRepo        string         `json:"giteeRepo,omitempty"`
	Branch           string         `json:"branch,omitempty"`
	GiteeAccessToken string         `json:"accessToken"` //access_token
	PodSpec          corev1.PodSpec `json:"podSpec"`
}

// RepoScannerStatus defines the observed state of RepoScanner
type RepoScannerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// RepoScanner is the Schema for the reposcanners API
type RepoScanner struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RepoScannerSpec   `json:"spec,omitempty"`
	Status RepoScannerStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// RepoScannerList contains a list of RepoScanner
type RepoScannerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RepoScanner `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RepoScanner{}, &RepoScannerList{})
}
