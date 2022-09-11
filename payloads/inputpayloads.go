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
