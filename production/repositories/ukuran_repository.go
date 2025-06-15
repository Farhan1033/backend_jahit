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

func UpdateMeasurement(id uint, input *models.Measurement) error {
	var measurement models.Measurement

	if err := config.DB.First(&measurement, "id = ?", id).Error; err != nil {
		return err
	}

	measurement.Type = input.Type
	measurement.Chest = input.Chest
	measurement.Waist = input.Waist
	measurement.Hips = input.Hips
	measurement.Shoulder = input.Shoulder
	measurement.Length = input.Length
	measurement.Sleeve = input.Sleeve
	measurement.Note = input.Note
	measurement.UpdatedAt = input.UpdatedAt

	return config.DB.Save(&measurement).Error
}

func DeleteMeasurement(id uint) error {
	return config.DB.Delete(&models.Measurement{}, "id = ?", id).Error
}
