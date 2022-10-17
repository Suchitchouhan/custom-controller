package main

import (
	// "context"

	"kubecontroller/controllers"
	"kubecontroller/kcontrollers"
	"kubecontroller/models"
	"os"

	"kubecontroller/cronjobs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	// metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//docker run --name=kubecontrollermysql -e MYSQL_DATABASE=kubecontrollerdb -e MYSQL_USER=kubeuser -e MYSQL_PASSWORD=xrock@971768 -e MYSQL_RANDOM_ROOT_PASSWORD=yes -d mysql:latest
//docker run --name=kubecontrollermysql -e MYSQL_DATABASE=kubecontrollerdb -e MYSQL_USER=kubeuser -e MYSQL_PASSWORD=xrock@971768 -e MYSQL_ROOT_PASSWORD=xrock -p 3306:3306 -d mysql:latest

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		panic("Error loading .env file")
	}
	kcontrollers.SetupConfig(os.Getenv("KUBECONFIG"))
	models.ConnectDatabase(os.Getenv("DATABASE"), os.Getenv("DBHOST"), os.Getenv("DBPORT"), os.Getenv("DBUSERNAME"), os.Getenv("DBPASSWORD"))
	cronjobs.RunCronjob()

	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/GetNodesAPI", controllers.GetNodesAPI)
	r.POST("/GetNodesDetailsAPI", controllers.GetNodesDetailsAPI)
	r.GET("/GetNameSpaceAPI", controllers.GetNameSpaceAPI)
	r.POST("/CreateNameSpaceAPI", controllers.CreateNameSpaceAPI)
	r.POST("/GetPodListAPI", controllers.GetPodListAPI)
	r.POST("/GetPodDetailsAPI", controllers.GetPodDetailsAPI)
	r.POST("/GetPodLogsAPI", controllers.GetPodLogsAPI)
	r.POST("/GetDeploymentsAPI", controllers.GetDeploymentsAPI)
	r.POST("/GetDeploymentDetailsAPI", controllers.GetDeploymentDetailsAPI)
	r.POST("/GetReplicasetsAPI", controllers.GetReplicasetsAPI)
	r.POST("/GetReplicasetDetailsAPI", controllers.GetReplicasetDetailsAPI)
	r.POST("/GetstatefulSetsAPI", controllers.GetstatefulSetsAPI)
	r.POST("/GetstatefulSetDetailsAPI", controllers.GetstatefulSetDetailsAPI)
	r.POST("/GetIngressAPI", controllers.GetIngressAPI)
	r.POST("/GetIngressClassAPI", controllers.GetIngressClassAPI)
	r.POST("/GetServicesAPI", controllers.GetServicesAPI)
	r.POST("/GetServiceDetailsAPI", controllers.GetServiceDetailsAPI)
	r.POST("/GetPodMetriceAPI", controllers.GetPodMetriceAPI)
	r.GET("/GetNodeMetrics", controllers.GetNodeMetrics)
	//r.GET("/GetNodeMetricesResourceAPI", controllers.GetNodeMetricesResourceAPI)
	r.Run(os.Getenv("HOST") + ":" + os.Getenv("PORT"))

}

//http://localhost:8001/api/v1/nodes
//http://localhost:8001/api/v1/pods
//http://localhost:8001/api/v1/services
///apis/extensions/v1beta1/namespaces/<namespace>/deployments
