package models

type IngredientePizza struct {
	IngredientePizzaId   int `json:"ingrediente_pizza_id" gorm:"unique;primaryKey;autoIncrement"`
	IngredienteId        int `json:"ingrediente_id"`
	PizzaPersonalizadaId int `json:"pizza_personalizada_id"`
}

func NewIngredientePizza(ingrediente, pizzaPersonalizada int) *IngredientePizza {
	return &IngredientePizza{
		IngredienteId:        ingrediente,
		PizzaPersonalizadaId: pizzaPersonalizada,
	}
}
