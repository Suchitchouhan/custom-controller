package controllers

import (
	"kubecontroller/kcontrollers"
	"kubecontroller/payloads"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetReplicasetsAPI(c *gin.Context) {
	var NameSpaceInputPayload payloads.NameSpaceInputPayload
	if err := c.ShouldBindJSON(&NameSpaceInputPayload); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	NS, err := kcontrollers.GetReplicaset(NameSpaceInputPayload.NameSpace)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Failed to Get ",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": NS})
}

func GetReplicasetDetailsAPI(c *gin.Context) {
	var ReplicaSetInputPayload payloads.ReplicaSetInputPayload
	if err := c.ShouldBindJSON(&ReplicaSetInputPayload); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	NS, err := kcontrollers.GetReplicasetDetails(ReplicaSetInputPayload.NameSpace, ReplicaSetInputPayload.ReplicaName)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Failed to Get ",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": NS})
}
