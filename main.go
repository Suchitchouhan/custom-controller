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
	r.GET("/GetNodesAPI", controllers.GetNodesAPI)
	r.POST("/GetNodesDetailsAPI", controllers.GetNodesDetailsAPI)
	r.GET("/GetNameSpaceAPI", controllers.GetNameSpaceAPI)
	r.POST("/CreateNameSpaceAPI", controllers.CreateNameSpaceAPI)
	r.POST("/GetPodListAPI", controllers.GetPodListAPI)
	r.POST("/GetPodDetailsAPI", controllers.GetPodDetailsAPI)
	r.POST("/GetPodLogsAPI", controllers.GetPodLogsAPI)
	r.Run("0.0.0.0:8000")

}

//http://localhost:8001/api/v1/nodes
//http://localhost:8001/api/v1/pods
//http://localhost:8001/api/v1/services
///apis/extensions/v1beta1/namespaces/<namespace>/deployments
