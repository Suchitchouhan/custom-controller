package kcontrollers

import (
	"flag"
	"fmt"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var ClientSet *kubernetes.Clientset

func SetupConfig() {
	var kubeconfig *string
	kubeconfig = flag.String("kubeconfig", "/home/suchit/.kube/config", "(optional) absolute path to the kubeconfig file")
	fmt.Println(*kubeconfig)
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	// fmt.Println(config)
	cltset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	ClientSet = cltset

}
