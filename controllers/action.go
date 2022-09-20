package controllers

import (
	"kubecontroller/kcontrollers"
	"kubecontroller/payloads"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetNameSpaceAPI(c *gin.Context) {
	GetData, err := kcontrollers.GetNameSpace()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Failed to upload",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Data": GetData, "Total NameSpace": len(GetData), "Message": "success"})

}

func CreateNameSpaceAPI(c *gin.Context) {
	var PodInputPayload payloads.NameSpaceInputPayload
	if err := c.ShouldBindJSON(&PodInputPayload); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	NS, err := kcontrollers.CreateNameSpace(PodInputPayload.NameSpace)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Failed to Create",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": NS})

}

func GetPodListAPI(c *gin.Context) {
	var PodInputPayload payloads.NameSpaceInputPayload
	if err := c.ShouldBindJSON(&PodInputPayload); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pods, err := kcontrollers.GetPodList(PodInputPayload.NameSpace)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "Pods": pods, "Total_pods": len(pods)})

}

func GetPodDetailsAPI(c *gin.Context) {
	var PodInputPayload payloads.PodInputPayload
	if err := c.ShouldBindJSON(&PodInputPayload); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := kcontrollers.GetPodDetails(PodInputPayload.NameSpace, PodInputPayload.PodName)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Message": data})

}

func GetPodLogsAPI(c *gin.Context) {
	var PodInputPayload payloads.PodInputPayload
	if err := c.ShouldBindJSON(&PodInputPayload); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := kcontrollers.GetPodLogs(PodInputPayload.NameSpace, PodInputPayload.PodName)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"log": data})

}

func GetPodMetriceAPI(c *gin.Context) {
	var InputPayload payloads.NameSpaceInputPayload
	if err := c.ShouldBindJSON(&InputPayload); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	NS, err := kcontrollers.GetPodMetrice(InputPayload.NameSpace)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Failed to Create",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": NS})
}

func GetNodesAPI(c *gin.Context) {
	data, err := kcontrollers.GetNodes()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Nodes": data})

}

func GetNodesDetailsAPI(c *gin.Context) {
	var PodInputPayload payloads.NodeNameInputPayload
	if err := c.ShouldBindJSON(&PodInputPayload); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := kcontrollers.GetNodesDetails(PodInputPayload.NodeName)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Node": data})
}

func GetNodeMetrics(c *gin.Context) {
	data, err := kcontrollers.GetNodeMetricesResource()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Node": data})

}
