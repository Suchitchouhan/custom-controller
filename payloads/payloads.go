package payloads

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/metrics/pkg/apis/metrics/v1beta1"
)

type KNameSpace struct {
	Name   string `json:"Name"`
	Status string `json:"Status"`
}

type KPod struct {
	UID               string            `json:"UID"`
	Name              string            `json:"Name"`
	Namespace         string            `json:"Namespace"`
	Age               string            `json:"Age"`
	RestartCount      string            `json:"RestartCount"`
	Status            string            `json:"Status"`
	APIVersion        string            `json:"APIVersion"`
	CreationTimestamp string            `json:"CreationTimestamp"`
	Labels            map[string]string `json:"labels"`
	HostIP            string            `json:"HostIP"`
	PodIP             string            `json:"PodIP"`
	Message           string            `json:"Message"`
}

type KNode struct {
	Name                    string               `json:"name"`
	UID                     string               `json:"uid"`
	ResourceVersion         string               `json:"resourceVersion"`
	CreationTimestamp       metav1.Time          `json:"creationTimestamp"`
	Labels                  map[string]string    `json:"labels"`
	MachineID               string               `json:"machineID"`
	SystemUUID              string               `json:"systemUUID"`
	BootID                  string               `json:"bootID"`
	KernelVersion           string               `json:"kernelVersion"`
	OsImage                 string               `json:"osImage"`
	ContainerRuntimeVersion string               `json:"containerRuntimeVersion"`
	KubeletVersion          string               `json:"kubeletVersion"`
	KubeProxyVersion        string               `json:"kubeProxyVersion"`
	OperatingSystem         string               `json:"operatingSystem"`
	Architecture            string               `json:"architecture"`
	Addresses               []corev1.NodeAddress `json:"addresses"`
}

type Kdeployment struct {
	Name               string            `json:"name"`
	Namespace          string            `json:"namespace"`
	UID                string            `json:"uid"`
	ResourceVersion    string            `json:"resourceVersion"`
	Generation         int               `json:"generation"`
	CreationTimestamp  metav1.Time       `json:"creationTimestamp"`
	Labels             map[string]string `json:"labels"`
	ObservedGeneration int               `json:"observedGeneration"`
	Replicas           int               `json:"replicas"`
	UpdatedReplicas    int               `json:"updatedReplicas"`
	ReadyReplicas      int               `json:"readyReplicas"`
	AvailableReplicas  int               `json:"availableReplicas"`
}

type KReplicaset struct {
	Name                 string            `json:"name"`
	Namespace            string            `json:"namespace"`
	UID                  string            `json:"uid"`
	ResourceVersion      string            `json:"resourceVersion"`
	Generation           int               `json:"generation"`
	CreationTimestamp    metav1.Time       `json:"creationTimestamp"`
	Labels               map[string]string `json:"labels"`
	Replicas             int               `json:"replicas"`
	FullyLabeledReplicas int               `json:"fullyLabeledReplicas"`
	ReadyReplicas        int               `json:"readyReplicas"`
	AvailableReplicas    int               `json:"availableReplicas"`
	ObservedGeneration   int               `json:"observedGeneration"`
}

type KStateFulSets struct {
	Name               string            `json:"name"`
	Namespace          string            `json:"namespace"`
	UID                string            `json:"uid"`
	ResourceVersion    string            `json:"resourceVersion"`
	Generation         int               `json:"generation"`
	CreationTimestamp  metav1.Time       `json:"creationTimestamp"`
	Labels             map[string]string `json:"labels"`
	ObservedGeneration int               `json:"observedGeneration"`
	Replicas           int               `json:"replicas"`
	ReadyReplicas      int               `json:"readyReplicas"`
	CurrentReplicas    int               `json:"currentReplicas"`
	UpdatedReplicas    int               `json:"updatedReplicas"`
	CurrentRevision    string            `json:"currentRevision"`
	UpdateRevision     string            `json:"updateRevision"`
	CollisionCount     int               `json:"collisionCount"`
	AvailableReplicas  int               `json:"availableReplicas"`
}

type KServices struct {
	Name              string             `json:"name"`
	Namespace         string             `json:"namespace"`
	UID               string             `json:"uid"`
	ResourceVersion   string             `json:"resourceVersion"`
	Generation        int                `json:"generation"`
	CreationTimestamp metav1.Time        `json:"creationTimestamp"`
	Labels            map[string]string  `json:"labels"`
	Spec              corev1.ServiceSpec `json:"spec"`
}

type KPodMetric struct {
	Pod         v1beta1.PodMetrics `json:"pod"`
	TotalMemory int                `json:"totalmemory"`
	TotalCpu    float64            `json:"totalcpu"`
}
