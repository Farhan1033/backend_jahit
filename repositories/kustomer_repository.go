package repositories

import (
	"sistem-ukuran-jahit/config"
	"sistem-ukuran-jahit/models"
	"time"

	"github.com/google/uuid"
)

func CreateCustomer(customer *models.Customer) error {
	return config.DB.Create(customer).Error
}

func GetAllCustomer() ([]models.Customer, error) {
	var customer []models.Customer
	err := config.DB.Find(&customer).Error
	return customer, err
}

func UpdateCustomer(id uuid.UUID, input *models.Customer) error {
	var customer models.Customer

	if err := config.DB.First(&customer, "id = ?", id).Error; err != nil {
		return err
	}

	customer.Name = input.Name
	customer.Phone = input.Phone
	customer.CreatedAt = time.Now()

	return config.DB.Save(&customer).Error
}

func DeleteCustomer(id uuid.UUID) error {
	return config.DB.Delete(&models.Customer{}, "id = ?", id).Error
}
