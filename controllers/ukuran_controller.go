package controllers

import (
	"net/http"
	"sistem-ukuran-jahit/config"
	"sistem-ukuran-jahit/models"
	"sistem-ukuran-jahit/models/dto"
	"sistem-ukuran-jahit/repositories"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetMeasurements(c *gin.Context) {
	var measurements []models.Measurement

	err := config.DB.Preload("Customer").Find(&measurements).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	type ResponseData struct {
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}

	var response []dto.MeasurementResponse

	for _, m := range measurements {
		r := dto.MeasurementResponse{
			ID:        m.ID,
			Name:      m.Customer.Name,
			Phone:     m.Customer.Phone,
			Type:      m.Type,
			Chest:     m.Chest,
			Waist:     m.Waist,
			Hips:      m.Hips,
			Shoulder:  m.Shoulder,
			Length:    m.Length,
			Sleeve:    m.Sleeve,
			Note:      m.Note,
			CreatedAt: m.CreatedAt.Format("2006-01-02 15:04"),
			UpdatedAt: m.UpdatedAt.Format("2006-01-02 15:04"),
		}
		response = append(response, r)
	}

	c.JSON(http.StatusOK, ResponseData{
		Message: "Berhasil menampilkan data",
		Data:    response,
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

	id := c.Param("id")
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
		Type      string  `json:"type"`
		Chest     float32 `json:"chest"`
		Waist     float32 `json:"waist"`
		Hips      float32 `json:"hips"`
		Shoulder  float32 `json:"shoulder"`
		Length    float32 `json:"length"`
		Sleeve    float32 `json:"sleeve"`
		Note      string  `json:"note"`
		UpdatedAt string  `json:"updated_at"`
	}

	type SuccessResponse struct {
		Message string              `json:"message"`
		Data    MeasurementResponse `json:"data"`
	}

	response := SuccessResponse{
		Message: "Berhasil memperbarui data",
		Data: MeasurementResponse{
			Type:      input.Type,
			Chest:     input.Chest,
			Waist:     input.Waist,
			Hips:      input.Hips,
			Shoulder:  input.Shoulder,
			Length:    input.Length,
			Sleeve:    input.Sleeve,
			Note:      input.Note,
			UpdatedAt: time.Now().Format(time.RFC3339),
		},
	}

	c.JSON(http.StatusOK, response)
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func DeleteMeasurement(c *gin.Context) {
	id := c.Param("id")
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

func GetMeasurementsByCustomerName(c *gin.Context) {
	name := c.Query("name")

	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter 'name' wajib diisi"})
		return
	}

	var measurements []models.Measurement

	err := config.DB.
		Joins("JOIN customers ON customers.id = measurements.customer_id").
		Where("LOWER(customers.name) LIKE LOWER(?)", "%"+name+"%").
		Preload("Customer").
		Find(&measurements).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response []dto.MeasurementResponse
	for _, m := range measurements {
		r := dto.MeasurementResponse{
			ID:        m.ID,
			Name:      m.Customer.Name,
			Phone:     m.Customer.Phone,
			Type:      m.Type,
			Chest:     m.Chest,
			Waist:     m.Waist,
			Hips:      m.Hips,
			Shoulder:  m.Shoulder,
			Length:    m.Length,
			Sleeve:    m.Sleeve,
			Note:      m.Note,
			CreatedAt: m.CreatedAt.Format("2006-01-02 15:04"),
			UpdatedAt: m.UpdatedAt.Format("2006-01-02 15:04"),
		}
		response = append(response, r)
	}

	type ResponseData struct {
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}

	c.JSON(http.StatusOK, ResponseData{
		Message: "Berhasil menampilkan data",
		Data: response,
	})
}
