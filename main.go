package main

import (
	// "context"

	"kubecontroller/controllers"
	"kubecontroller/kcontrollers"

	"github.com/gin-gonic/gin"
	// metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	kcontrollers.SetupConfig()

	r := gin.Default()
	r.Use(gin.Logger())
	r.GET("/GetNameSpaceAPI", controllers.GetNameSpaceAPI)
	r.POST("/CreateNameSpaceAPI", controllers.CreateNameSpaceAPI)
	r.POST("/GetPodListAPI", controllers.GetPodListAPI)
	r.Run("0.0.0.0:8000")

}
