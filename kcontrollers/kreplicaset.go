package kcontrollers

import (
	"context"
	"kubecontroller/payloads"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetReplicaset(NameSpace string) ([]payloads.KReplicaset, error) {
	RS, err := ClientSet.AppsV1().ReplicaSets(NameSpace).List(context.Background(), metav1.ListOptions{})
	var final_KReplicaset []payloads.KReplicaset
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
	if err != nil {
		return final_KReplicaset, err
	}
	return final_KReplicaset, nil
}
