package controllers

import (
	"kubecontroller/kcontrollers"
	"kubecontroller/payloads"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDeploymentsAPI(c *gin.Context) {
	var PodInputPayload payloads.NameSpaceInputPayload
	if err := c.ShouldBindJSON(&PodInputPayload); err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	NS, err := kcontrollers.GetDeployments(PodInputPayload.NameSpace)
	if err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err,
			"message": "Failed to Get ",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": NS})
}

func GetDeploymentDetailsAPI(c *gin.Context) {
	var DeploymentInputPayload payloads.DeploymentInputPayload
	if err := c.ShouldBindJSON(&DeploymentInputPayload); err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	NS, err := kcontrollers.GetDeploymentDetails(DeploymentInputPayload.NameSpace, DeploymentInputPayload.DeploymentName)
	if err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err,
			"message": "Failed to Get",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": NS})
}
