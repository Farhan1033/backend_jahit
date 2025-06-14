package routes

import (
	"sistem-ukuran-jahit/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/measurements", controllers.GetMeasurements)
		api.POST("/measurements", controllers.CreateMeasurement)
		api.PUT("/measurements/:id_customer", controllers.UpdateMeasurement)
		api.DELETE("/measurements/:id_customer", controllers.DeleteMeasurement)
	}
}
