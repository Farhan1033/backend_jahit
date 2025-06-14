package main

import (
	"sistem-ukuran-jahit/config"
	"sistem-ukuran-jahit/models"
	"sistem-ukuran-jahit/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.Measurement{})

	r := gin.Default()
	routes.SetupRoutes(r)

	r.Run(":8080")
}
