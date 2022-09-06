package kcontrollers

import (
	"context"
	"fmt"
	"kubecontroller/payloads"
	"log"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetNameSpace() ([]payloads.KNameSpace, error) {
	nsList, err := ClientSet.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
	var Namespace []payloads.KNameSpace
	if err != nil {
		log.Fatalf("Error : %v \n", err)
		return Namespace, err
	}
	for _, n := range nsList.Items {
		ns := payloads.KNameSpace{Name: fmt.Sprint(n.Name), Status: fmt.Sprint(n.Status.Phase)}
		Namespace = append(Namespace, ns)
	}
	return Namespace, nil
}

func CreateNameSpace(Name string) (string, error) {
	nsName := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: Name,
		},
	}
	__, err := ClientSet.CoreV1().Namespaces().Create(context.Background(), nsName, metav1.CreateOptions{})
	if err != nil {
		log.Fatalf("Error : %v \n", err)
		return "fail", err
	}
	return "success", nil
}
