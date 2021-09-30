package cd

import (
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/printers"
	cliresource "k8s.io/cli-runtime/pkg/resource"
	"k8s.io/kubectl/pkg/cmd/apply"
	"k8s.io/kubectl/pkg/cmd/delete"
	"log"
	"os"
)

func DeployWebUI() {

	var (
		kubeConfigFile string = os.Getenv("HOME") + "/.kube/config"
		err            error
	)
	kubeConfigFlags := genericclioptions.NewConfigFlags(true).WithDeprecatedPasswordFlag()
	kubeConfigFlags.KubeConfig = &kubeConfigFile

	builder := cliresource.NewBuilder(kubeConfigFlags)
	if err != nil {
		log.Println(err)
		return
	}
	ioStreams := genericclioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr}
	o := apply.NewApplyOptions(ioStreams)
	// default namespace for k8s resource
	o.Namespace = "default"
	o.Builder = builder
	o.DeleteOptions = &delete.DeleteOptions{
		FilenameOptions: cliresource.FilenameOptions{
			// target k8s yaml files and directories that contain k8s yaml files
			Filenames:[]string{"/opt"},

			Recursive: false,
		},
	}
	o.ToPrinter = func(operation string) (printers.ResourcePrinter, error) {
		o.PrintFlags.NamePrintFlags.Operation = operation
		return o.PrintFlags.ToPrinter()
	}


	err = o.Run()
	if err != err {
		log.Println(err)
	}
}

func RemoveWebUI(name string){

var (
		kubeConfigFile string = os.Getenv("HOME") + "/.kube/config"
		err            error
	)
	kubeConfigFlags := genericclioptions.NewConfigFlags(true).WithDeprecatedPasswordFlag()
	kubeConfigFlags.KubeConfig = &kubeConfigFile

	builder := cliresource.NewBuilder(kubeConfigFlags)
	if err != nil {
		log.Println(err)
		return
	}
	ioStreams := genericclioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr}
	o := apply.DeleteOptions(ioStreams)
	// default namespace for k8s resource
	o.Namespace = "default"
	o.Builder = builder
	o.DeleteOptions = &delete.DeleteOptions{
		FilenameOptions: cliresource.FilenameOptions{
			// target k8s yaml files and directories that contain k8s yaml files
			Filenames:[]string{"/opt"},

			Recursive: false,
		},
	}
	o.ToPrinter = func(operation string) (printers.ResourcePrinter, error) {
		o.PrintFlags.NamePrintFlags.Operation = operation
		return o.PrintFlags.ToPrinter()
	}


	err = o.Run()
	if err != err {
		log.Println(err)
	}
}