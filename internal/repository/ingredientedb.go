package repository

import (
	"github.com/carlosclavijo/pizza/internal/database"
	"github.com/carlosclavijo/pizza/internal/models"
)

func InsertIngrediente(ingrediente *models.Ingrediente) error {
	db := database.GetInstance()
	result := db.Db.Create(&ingrediente)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetIngredientes(ingredientes *[]models.Ingrediente) error {
	db := database.GetInstance()
	result := db.Db.Find(&ingredientes)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetIngredientById(ingrediente *models.Ingrediente, id int) error {
	db := database.GetInstance()
	result := db.Db.First(&ingrediente, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
