package repository

import (
	"github.com/carlosclavijo/pizza/internal/database"
	"github.com/carlosclavijo/pizza/internal/models"
)

func InsertPedido(pedido *models.Pedido) error {
	db := database.GetInstance()
	result := db.Db.Create(&pedido)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
