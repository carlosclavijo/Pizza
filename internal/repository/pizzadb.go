package repository

import (
	"github.com/carlosclavijo/pizza/internal/database"
	"github.com/carlosclavijo/pizza/internal/models"
)

func InsertarPizza(pizza models.IPizza) error {
	db := database.GetInstance()
	tipo := pizza.GetTipo()
	if tipo == "PizzaSabor" {
		var pizzasabor *models.PizzaSabor = pizza.(*models.PizzaSabor)
		result := db.Db.Create(pizzasabor)
		if result.Error != nil {
			return result.Error
		}
		result = db.Db.First(&pizzasabor.Sabor, pizzasabor.SaborId)
		if result.Error != nil {
			return result.Error
		}
	} else {
		var pizzapersonalizada *models.PizzaPersonalizada = pizza.(*models.PizzaPersonalizada)
		result := db.Db.Create(&pizzapersonalizada)
		if result.Error != nil {
			return result.Error
		}
		for _, s := range pizzapersonalizada.Ingredientes {
			ingredientepiza := models.NewIngredientePizza(s.IngredienteId, pizzapersonalizada.PizzaPersonalizadaId)
			result := db.Db.Create(&ingredientepiza)
			if result.Error != nil {
				return result.Error
			}
		}

	}
	return nil
}
