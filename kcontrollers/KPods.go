package kcontrollers

import (
	"context"
	"fmt"
	"kubecontroller/payloads"
	"time"

	// corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	// "k8s.io/client-go/kubernetes"
)

func Age(pod *v1.Pod) uint {
	diff := uint(time.Now().Sub(pod.Status.StartTime.Time).Hours())
	return uint(diff / 24)
}

func RestartCount(pod *v1.Pod) int32 {
	if len(pod.Status.ContainerStatuses) > 0 {
		return pod.Status.ContainerStatuses[0].RestartCount
	}

	return 0
}

func Status(pod *v1.Pod) v1.PodPhase {
	return pod.Status.Phase
}

func GetPodList(namespace string) ([]payloads.KPod, error) {
	var kpod []payloads.KPod
	pods, err := ClientSet.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	// fmt.Println(pods.Items)
	for _, value := range pods.Items {

		// data, _ := json.Marshal(value)
		// fmt.Println("data")
		// fmt.Println(string(data))
		pk := payloads.KPod{}
		pk.UID = fmt.Sprint(value.UID)
		pk.Name = fmt.Sprint(value.Name)
		pk.Namespace = fmt.Sprint(value.Namespace)
		pk.Age = fmt.Sprint(Age(&value))
		pk.RestartCount = fmt.Sprint(RestartCount(&value))
		pk.Status = fmt.Sprint(Status(&value))
		pk.APIVersion = fmt.Sprint(value.APIVersion)
		pk.CreationTimestamp = fmt.Sprint(value.CreationTimestamp)
		pk.HostIP = fmt.Sprint(value.Status.HostIP)
		pk.PodIP = fmt.Sprint(value.Status.PodIP)
		pk.Message = fmt.Sprint(value.Status.Message)
		kpod = append(kpod, pk)
	}
	if err != nil {
		return kpod, err
	}

	return kpod, nil
}

// func GetDetails() {

// }
