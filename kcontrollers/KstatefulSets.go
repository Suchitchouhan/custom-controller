package kcontrollers

import (
	"context"
	"kubecontroller/payloads"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetstatefulSet(NameSpace string) ([]payloads.KStateFulSets, error) {
	var KStateFulSets []payloads.KStateFulSets
	SS, err := ClientSet.AppsV1().StatefulSets(NameSpace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return KStateFulSets, err
	}
	for _, value := range SS.Items {
		KSF := payloads.KStateFulSets{}
		KSF.Name = value.Name
		KSF.Namespace = value.Namespace
		KSF.UID = string(value.UID)
		KSF.Generation = int(value.Generation)
		KSF.Labels = value.Labels
		KSF.ObservedGeneration = int(value.Status.ObservedGeneration)
		KSF.Replicas = int(value.Status.Replicas)
		KSF.ReadyReplicas = int(value.Status.ReadyReplicas)
		KSF.CurrentReplicas = int(value.Status.CurrentReplicas)
		KSF.UpdatedReplicas = int(value.Status.UpdatedReplicas)
		KSF.CurrentRevision = value.Status.CurrentRevision
		KSF.UpdateRevision = value.Status.UpdateRevision
		KSF.CollisionCount = int(*value.Status.CollisionCount)
		KSF.AvailableReplicas = int(value.Status.AvailableReplicas)
		KStateFulSets = append(KStateFulSets, KSF)
	}
	return KStateFulSets, err
}

func GetstatefulSetDetails(NameSpace string, StatefulName string) (*appsv1.StatefulSet, error) {
	SS, err := ClientSet.AppsV1().StatefulSets(NameSpace).Get(context.TODO(), StatefulName, metav1.GetOptions{})
	if err != nil {
		return &appsv1.StatefulSet{}, err
	}
	return SS, nil
}
