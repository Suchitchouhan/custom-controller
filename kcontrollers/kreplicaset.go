package kcontrollers

import (
	"context"
	"kubecontroller/payloads"
	"log"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetReplicaset(NameSpace string) ([]payloads.KReplicaset, error) {
	RS, err := ClientSet.AppsV1().ReplicaSets(NameSpace).List(context.Background(), metav1.ListOptions{})
	var final_KReplicaset []payloads.KReplicaset
	if err != nil {
		log.Fatalf("Error : %v \n", err)
		return final_KReplicaset, err
	}
	for _, value := range RS.Items {

		prk := payloads.KReplicaset{}
		prk.Name = value.Name
		prk.Namespace = value.Namespace
		prk.UID = string(value.UID)
		prk.ResourceVersion = value.ResourceVersion
		prk.Generation = int(value.Generation)
		prk.CreationTimestamp = value.CreationTimestamp
		prk.Labels = value.Labels
		prk.Replicas = int(value.Status.Replicas)
		prk.FullyLabeledReplicas = int(value.Status.FullyLabeledReplicas)
		prk.ReadyReplicas = int(value.Status.ReadyReplicas)
		prk.AvailableReplicas = int(value.Status.AvailableReplicas)
		prk.ObservedGeneration = int(value.Status.ObservedGeneration)
		final_KReplicaset = append(final_KReplicaset, prk)
	}

	return final_KReplicaset, nil
}

func GetReplicasetDetails(NameSpace string, ReplicaName string) (*appsv1.ReplicaSet, error) {
	RS, err := ClientSet.AppsV1().ReplicaSets(NameSpace).Get(context.Background(), ReplicaName, metav1.GetOptions{})
	if err != nil {
		log.Fatalf("Error : %v \n", err)
		return &appsv1.ReplicaSet{}, err
	}
	return RS, nil
}
