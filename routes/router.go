package routes

import (
	"sistem-ukuran-jahit/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/measurements", controllers.GetMeasurements)
		api.GET("/measurement/customer", controllers.GetMeasurementsByCustomerName)
		api.POST("/measurements", controllers.CreateMeasurement)
		api.PUT("/measurements/:id", controllers.UpdateMeasurement)
		api.DELETE("/measurements/:id", controllers.DeleteMeasurement)
		api.POST("/add-customer", controllers.CreateCustomer)
		api.GET("/customer", controllers.GetAllCustomer)
		api.DELETE("/delete-customer", controllers.DeleteCustomer)
	}
}
