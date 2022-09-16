package kcontrollers

import (
	"context"
	"kubecontroller/payloads"

	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetIngress(NameSpace string) (*v1.IngressList, error) {
	IN, err := ClientSet.NetworkingV1().Ingresses(NameSpace).List(context.Background(), metav1.ListOptions{})
	return IN, err
}

func GetIngressClass(NameSpace string) (*v1.IngressClassList, error) {
	IN, err := ClientSet.NetworkingV1().IngressClasses().List(context.Background(), metav1.ListOptions{})
	return IN, err
}

func GetServices(NameSpace string) ([]payloads.KServices, error) {
	var KServices []payloads.KServices
	IN, err := ClientSet.CoreV1().Services(NameSpace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return KServices, err
	}
	for _, value := range IN.Items {
		KS := payloads.KServices{}
		KS.Name = value.Name
		KS.Namespace = value.Namespace
		KS.UID = string(value.UID)
		KS.ResourceVersion = value.ResourceVersion
		KS.CreationTimestamp = value.CreationTimestamp
		KS.Labels = value.Labels
		KS.Spec = value.Spec
		KServices = append(KServices, KS)

	}
	return KServices, nil
}

func GetServiceDetails(NameSpace string, ServiceName string) (*corev1.Service, error) {
	IN, err := ClientSet.CoreV1().Services(NameSpace).Get(context.Background(), ServiceName, metav1.GetOptions{})
	return IN, err
}
