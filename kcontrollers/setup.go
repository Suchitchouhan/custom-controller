package kcontrollers

import (
	"flag"
	"fmt"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

//var MC *metrics.Clientset

var ClientSet *kubernetes.Clientset

func SetupConfig(kubeconfigpath string) {
	var kubeconfig *string
	kubeconfig = flag.String("kubeconfig", kubeconfigpath, "(optional) absolute path to the kubeconfig file")
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

	//mclientset, err := metrics.NewForConfig(config)

	//if err != nil {
	//	panic(err)
	//}

	//MC = mclientset
	ClientSet = cltset

}
