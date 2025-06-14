package controllers

import (
	"net/http"
	"sistem-ukuran-jahit/models"
	"sistem-ukuran-jahit/repositories"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetMeasurements(c *gin.Context) {
	data, err := repositories.GetAllMeasurements()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	type Response struct {
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}

	c.JSON(http.StatusOK, Response{
		Message: "Berhasil menampilkan data",
		Data:    data,
	})
}

func CreateMeasurement(c *gin.Context) {
	var measurement models.Measurement
	if err := c.ShouldBindJSON(&measurement); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := repositories.CreateMeasurement(&measurement); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	type Response struct {
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}

	c.JSON(http.StatusCreated, Response{
		Message: "Berhasil menambahkan data",
		Data:    measurement,
	})
}

func UpdateMeasurement(c *gin.Context) {
	var input models.Measurement

	id := c.Param("id_customer")
	parsedID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Format ID tidak valid"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	if err := repositories.UpdateMeasurement(uint(parsedID), &input); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	type MeasurementResponse struct {
		CustomerName string  `json:"customer_name"`
		Gender       string  `json:"gender"`
		Chest        float64 `json:"chest"`
		Waist        float64 `json:"waist"`
		Hip          float64 `json:"hip"`
		ArmLength    float64 `json:"arm_length"`
		ShirtLength  float64 `json:"shirt_length"`
		UpdatedAt    string  `json:"updated_at"`
	}

	type SuccessResponse struct {
		Message string              `json:"message"`
		Data    MeasurementResponse `json:"data"`
	}

	response := SuccessResponse{
		Message: "Berhasil memperbarui data",
		Data: MeasurementResponse{
			CustomerName: input.CustomerName,
			Gender:       input.Gender,
			Chest:        input.Chest,
			Waist:        input.Waist,
			Hip:          input.Hip,
			ArmLength:    input.ArmLength,
			ShirtLength:  input.ShirtLength,
			UpdatedAt:    time.Now().Format(time.RFC3339),
		},
	}

	c.JSON(http.StatusOK, response)
}

type ErrorResponse struct {
	Error string `json:"error"`
}



func DeleteMeasurement(c *gin.Context) {
	id := c.Param("id_customer")
	parsedID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := repositories.DeleteMeasurement(uint(parsedID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	type Response struct {
		Message string `json:"message"`
	}

	c.JSON(http.StatusOK, Response{
		Message: "Berhasil menghapus data",
	})
}
