package kcontrollers

import (
	"context"
	"fmt"

	"kubecontroller/payloads"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetDeployments(NameSpace string) ([]payloads.Kdeployment, error) {
	var Kdeployment []payloads.Kdeployment
	depl, err := ClientSet.AppsV1().Deployments(NameSpace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return Kdeployment, err
	}
	for _, value := range depl.Items {
		pk := payloads.Kdeployment{}
		pk.Name = value.Name
		pk.Namespace = value.Namespace
		pk.UID = fmt.Sprint(value.UID)
		pk.Generation = int(value.Generation)
		pk.CreationTimestamp = value.CreationTimestamp
		pk.Labels = value.Labels
		pk.ObservedGeneration = int(value.Status.ObservedGeneration)
		pk.Replicas = int(*value.Spec.Replicas)
		pk.UpdatedReplicas = int(value.Status.UpdatedReplicas)
		pk.ReadyReplicas = int(value.Status.ReadyReplicas)
		pk.AvailableReplicas = int(value.Status.AvailableReplicas)
		Kdeployment = append(Kdeployment, pk)

	}

	return Kdeployment, nil

}

func GetDeploymentDetails(NameSpace string, DeploymentName string) (*appsv1.Deployment, error) {
	depl, err := ClientSet.AppsV1().Deployments(NameSpace).Get(context.Background(), DeploymentName, metav1.GetOptions{})
	if err != nil {
		return &appsv1.Deployment{}, err
	}
	return depl, nil
}
