package main

import (
	"sistem-ukuran-jahit/config"
	"sistem-ukuran-jahit/models"
	"sistem-ukuran-jahit/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.Customer{}, &models.Measurement{})

	// gin.SetMode(gin.DebugMode)
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	routes.SetupRoutes(r)
	r.Run(":8080")
}
