package main

import (
	"fmt"
	"os"
	"sistem-ukuran-jahit/config"
	"sistem-ukuran-jahit/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	if err := r.SetTrustedProxies([]string{"0.0.0.0"}); err != nil {
		panic(fmt.Sprintf("Gagal set trusted proxy: %v", err))
	}

	r.GET("/kaithhealthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "OK"})
	})

	routes.SetupRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run("0.0.0.0:" + port)
}
