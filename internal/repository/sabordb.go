package repository

import (
	"github.com/carlosclavijo/pizza/internal/database"
	"github.com/carlosclavijo/pizza/internal/models"
)

func InsertSabor(sabor *models.Sabor) error {
	db := database.GetInstance()
	result := db.Db.Create(&sabor)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetSabores(sabor *[]models.Sabor) error {
	db := database.GetInstance()
	result := db.Db.Find(&sabor)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetSaborById(sabor *models.Sabor, id int) error {
	db := database.GetInstance()
	result := db.Db.First(&sabor, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
