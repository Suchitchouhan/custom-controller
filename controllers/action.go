package controllers

import (
	"log"
	"net/http"

	"kubecontroller/kcontrollers"

	"github.com/gin-gonic/gin"
)

func GetNameSpaceAPI(c *gin.Context) {
	GetData, err := kcontrollers.GetNameSpace()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err,
			"message": "Failed to upload",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Data": GetData, "Total NameSpace": len(GetData), "Message": "success"})
	return
}

func CreateNameSpaceAPI(c *gin.Context) {
	Name := c.PostForm("Name")
	log.Println(Name)
	if Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Name can not we empty",
			"message": "Failed to upload",
		})
		return
	}

	NS, err := kcontrollers.CreateNameSpace(Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err,
			"message": "Failed to Create",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": NS})
	return

}

func GetPodListAPI(c *gin.Context) {
	NameSpace := c.DefaultPostForm("NameSpace", "default")
	pods, err := kcontrollers.GetPodList(NameSpace)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err,
			"message": "Failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": pods})
	return

}
