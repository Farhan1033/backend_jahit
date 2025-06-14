package repositories

import (
	"sistem-ukuran-jahit/config"
	"sistem-ukuran-jahit/models"
)

func GetAllMeasurements() ([]models.Measurement, error) {
	var measurements []models.Measurement
	err := config.DB.Find(&measurements).Error
	return measurements, err
}

func CreateMeasurement(measurement *models.Measurement) error {
	return config.DB.Create(measurement).Error
}

func UpdateMeasurement(id_customer uint, input *models.Measurement) error {
	var measurement models.Measurement

	if err := config.DB.First(&measurement, "id_customer = ?", id_customer).Error; err != nil {
		return err
	}

	measurement.CustomerName = input.CustomerName
	measurement.Gender = input.Gender
	measurement.Chest = input.Chest
	measurement.Waist = input.Waist
	measurement.Hip = input.Hip
	measurement.ArmLength = input.ArmLength
	measurement.ShirtLength = input.ShirtLength
	measurement.UpdatedAt = input.UpdatedAt

	return config.DB.Save(&measurement).Error
}

func DeleteMeasurement(id_customer uint) error {
	return config.DB.Delete(&models.Measurement{}, "id_customer = ?", id_customer).Error
}
