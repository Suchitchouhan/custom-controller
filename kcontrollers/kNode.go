package kcontrollers

import (
	"context"
	"kubecontroller/payloads"
	"log"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/metrics/pkg/apis/metrics/v1beta1"
)

func GetNodes() ([]payloads.KNode, error) {
	var KNode []payloads.KNode
	nsList, err := ClientSet.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatalf("Error : %v \n", err)
		return KNode, err
	}
	for _, value := range nsList.Items {
		kn := payloads.KNode{}
		kn.Name = value.Name
		kn.UID = string(value.UID)
		kn.ResourceVersion = value.ResourceVersion
		kn.CreationTimestamp = value.CreationTimestamp
		kn.Labels = value.Labels
		kn.MachineID = value.Status.NodeInfo.MachineID
		kn.SystemUUID = value.Status.NodeInfo.SystemUUID
		kn.BootID = value.Status.NodeInfo.BootID
		kn.KernelVersion = value.Status.NodeInfo.KernelVersion
		kn.ContainerRuntimeVersion = value.Status.NodeInfo.ContainerRuntimeVersion
		kn.KubeletVersion = value.Status.NodeInfo.KubeletVersion
		kn.KubeProxyVersion = value.Status.NodeInfo.KubeProxyVersion
		kn.OperatingSystem = value.Status.NodeInfo.OperatingSystem
		kn.Architecture = value.Status.NodeInfo.Architecture
		kn.Addresses = value.Status.Addresses
		KNode = append(KNode, kn)
	}
	return KNode, nil
}

func GetNodesDetails(Nodename string) (*v1.Node, error) {
	ns, err := ClientSet.CoreV1().Nodes().Get(context.Background(), Nodename, metav1.GetOptions{})
	if err != nil {
		log.Fatalf("Error : %v \n", err)
		return &v1.Node{}, err
	}

	return ns, nil
}

func GetNodeMetricesResource() (*v1beta1.NodeMetricsList, error) {
	ns, err := MC.MetricsV1beta1().NodeMetricses().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatalf("Error : %v \n", err)
		return &v1beta1.NodeMetricsList{}, err
	}

	return ns, nil
}
