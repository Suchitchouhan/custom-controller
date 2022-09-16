package controllers

import (
	"kubecontroller/kcontrollers"
	"kubecontroller/payloads"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetIngressAPI(c *gin.Context) {
	var PodInputPayload payloads.NameSpaceInputPayload
	if err := c.ShouldBindJSON(&PodInputPayload); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	NS, err := kcontrollers.GetIngress(PodInputPayload.NameSpace)
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

func GetIngressClassAPI(c *gin.Context) {
	var PodInputPayload payloads.NameSpaceInputPayload
	if err := c.ShouldBindJSON(&PodInputPayload); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	NS, err := kcontrollers.GetIngressClass(PodInputPayload.NameSpace)
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

func GetServicesAPI(c *gin.Context) {
	var PodInputPayload payloads.NameSpaceInputPayload
	if err := c.ShouldBindJSON(&PodInputPayload); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	NS, err := kcontrollers.GetServices(PodInputPayload.NameSpace)
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

func GetServiceDetailsAPI(c *gin.Context) {
	var InputPayload payloads.ServiceInputPayload
	if err := c.ShouldBindJSON(&InputPayload); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	NS, err := kcontrollers.GetServiceDetails(InputPayload.NameSpace, InputPayload.ServiceName)
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
