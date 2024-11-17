package models

import "github.com/gofrs/uuid"

type PizzaPersonalizada struct {
	PizzaPersonalizadaId uuid.UUID
	Ingredientes         []Ingrediente
	Precio               float32
}

func (pizza PizzaPersonalizada) CalcularPrecio() float32 {
	var precio float32
	for _, ingrediente := range pizza.Ingredientes {
		precio += ingrediente.Precio
	}
	return precio
}

func NewPizzaPersonalizada(ingredientes []Ingrediente) iPizza {
	return &PizzaPersonalizada{
		Ingredientes: ingredientes,
	}
}
