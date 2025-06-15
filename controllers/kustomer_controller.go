package controllers

import (
	"net/http"
	"sistem-ukuran-jahit/models"
	"sistem-ukuran-jahit/repositories"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetAllCustomer(c *gin.Context) {
	data, err := repositories.GetAllCustomer()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	type responseCustomer struct {
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}

	c.JSON(http.StatusOK, responseCustomer{
		Message: "Berhasil menampilkan data",
		Data:    data,
	})
}

func CreateCustomer(c *gin.Context) {
	var customer models.Customer

	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer.ID = uuid.New()

	if err := repositories.CreateCustomer(&customer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	type Response struct {
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}

	c.JSON(http.StatusCreated, Response{
		Message: "Berhasil menambahkan data",
		Data:    customer,
	})
}
