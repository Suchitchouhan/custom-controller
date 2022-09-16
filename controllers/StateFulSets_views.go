package controllers

import (
	"kubecontroller/kcontrollers"
	"kubecontroller/payloads"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetstatefulSetsAPI(c *gin.Context) {
	var NameSpaceInputPayload payloads.NameSpaceInputPayload
	if err := c.ShouldBindJSON(&NameSpaceInputPayload); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	NS, err := kcontrollers.GetstatefulSet(NameSpaceInputPayload.NameSpace)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Failed to Get ",
		})
	}

	c.JSON(http.StatusOK, gin.H{"Message": NS})
}

func GetstatefulSetDetailsAPI(c *gin.Context) {
	var StateFulSetInputPayload payloads.StateFulSetInputPayload
	if err := c.ShouldBindJSON(&StateFulSetInputPayload); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	NS, err := kcontrollers.GetstatefulSetDetails(StateFulSetInputPayload.NameSpace, StateFulSetInputPayload.StatefulName)
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
