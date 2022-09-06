package payloads

type KNameSpace struct {
	Name   string `json:"Name"`
	Status string `json:"Status"`
}

type KPod struct {
	UID               string `json:"UID"`
	Name              string `json:"Name"`
	Namespace         string `json:"Namespace"`
	Age               string `json:"Age"`
	RestartCount      string `json:"RestartCount"`
	Status            string `json:"Status"`
	APIVersion        string `json:"APIVersion"`
	CreationTimestamp string `json:"CreationTimestamp"`
	HostIP            string `json:"HostIP"`
	PodIP             string `json:"PodIP"`
	Message           string `json:"Message"`
}
