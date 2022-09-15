package main

import (
	// "context"

	"kubecontroller/controllers"
	"kubecontroller/kcontrollers"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	// metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		panic("Error loading .env file")
	}
	kcontrollers.SetupConfig(os.Getenv("KUBECONFIG"))

	r := gin.Default()
	r.Use(gin.Logger())
	r.GET("/GetNodesAPI", controllers.GetNodesAPI)
	r.POST("/GetNodesDetailsAPI", controllers.GetNodesDetailsAPI)
	r.GET("/GetNameSpaceAPI", controllers.GetNameSpaceAPI)
	r.POST("/CreateNameSpaceAPI", controllers.CreateNameSpaceAPI)
	r.POST("/GetPodListAPI", controllers.GetPodListAPI)
	r.POST("/GetPodDetailsAPI", controllers.GetPodDetailsAPI)
	r.POST("/GetPodLogsAPI", controllers.GetPodLogsAPI)
	r.POST("/GetDeploymentsAPI", controllers.GetDeploymentsAPI)
	r.POST("/GetDeploymentDetailsAPI", controllers.GetDeploymentDetailsAPI)
	r.POST("/GetReplicasets", controllers.GetReplicasets)
	//r.GET("/GetNodeMetricesResourceAPI", controllers.GetNodeMetricesResourceAPI)
	r.Run(os.Getenv("HOST") + ":" + os.Getenv("PORT"))

}

//http://localhost:8001/api/v1/nodes
//http://localhost:8001/api/v1/pods
//http://localhost:8001/api/v1/services
///apis/extensions/v1beta1/namespaces/<namespace>/deployments
