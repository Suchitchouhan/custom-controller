package kcontrollers

import (
	"context"
	"log"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetNodes() (*v1.NodeList, error) {
	nsList, err := ClientSet.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatalf("Error : %v \n", err)
		return &v1.NodeList{}, err
	}

	return nsList, nil
}

func GetNodesDetails(Nodename string) (*v1.Node, error) {
	ns, err := ClientSet.CoreV1().Nodes().Get(context.Background(), Nodename, metav1.GetOptions{})
	if err != nil {
		log.Fatalf("Error : %v \n", err)
		return &v1.Node{}, err
	}

	return ns, nil
}



