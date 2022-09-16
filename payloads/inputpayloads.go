package payloads

type PodInputPayload struct {
	NameSpace string `json:"NameSpace" binding:"required"`
	PodName   string `json:"PodName" binding:"required"`
}

type NameSpaceInputPayload struct {
	NameSpace string `json:"NameSpace" binding:"required"`
}

type NodeNameInputPayload struct {
	NodeName string `json:"NodeName" binding:"required"`
}

type DeploymentInputPayload struct {
	NameSpace      string `json:"NameSpace" binding:"required"`
	DeploymentName string `json:"DeploymentName" binding:"required"`
}

type ReplicaSetInputPayload struct {
	NameSpace   string `json:"NameSpace" binding:"required"`
	ReplicaName string `json:"ReplicaName" binding:"required"`
}

type StateFulSetInputPayload struct {
	NameSpace    string `json:"NameSpace" binding:"required"`
	StatefulName string `json:"StatefulName" binding:"required"`
}

type ServiceInputPayload struct {
	NameSpace   string `json:"NameSpace" binding:"required"`
	ServiceName string `json:"ServiceName" binding:"required"`
}
