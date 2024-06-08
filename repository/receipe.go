package repository

import (
	"Ayala-Crea/ResepBe/model"

	"gorm.io/gorm"
)

func InsertReceipt(db *gorm.DB, receipe model.Receipt) error {
	if err := db.Create(&receipe).Error; err != nil {
		return err
	}
	return nil
}