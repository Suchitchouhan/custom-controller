package kcontrollers

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"kubecontroller/payloads"
	"time"

	// corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/metrics/pkg/apis/metrics/v1beta1"
	//metrics "k8s.io/metrics/pkg/client/clientset/versioned"
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
	if err != nil {
		return kpod, err
	}

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
		pk.Labels = value.Labels
		pk.HostIP = fmt.Sprint(value.Status.HostIP)
		pk.PodIP = fmt.Sprint(value.Status.PodIP)
		pk.Message = fmt.Sprint(value.Status.Message)
		kpod = append(kpod, pk)
	}

	return kpod, nil
}

func GetPodDetails(namespace string, podname string) (*v1.Pod, error) {

	pod, err := ClientSet.CoreV1().Pods(namespace).Get(context.Background(), podname, metav1.GetOptions{})
	if err != nil {
		return &v1.Pod{}, err
	}

	return pod, nil

}

func GetPodLogs(namespace string, podname string) (string, error) {
	podLogOpts := v1.PodLogOptions{}
	req := ClientSet.CoreV1().Pods(namespace).GetLogs(podname, &podLogOpts)
	podLogs, err := req.Stream(context.Background())
	if err != nil {
		return "error in opening stream", err
	}
	defer podLogs.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, podLogs)
	if err != nil {
		return "error in copy information from podLogs to buf", err
	}
	str := buf.String()

	return str, nil
}

func GetPodMetrice(namespace string) (*v1beta1.PodMetricsList, error) {
	PM, err := MC.MetricsV1beta1().PodMetricses(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return &v1beta1.PodMetricsList{}, err
	}
	return PM, err
}
