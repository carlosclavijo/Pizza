package models

import "github.com/gofrs/uuid"

type PizzaSabor struct {
	PizzaSaborId uuid.UUID
	Nombre       string
	Precio       float32
}

func (pizza PizzaSabor) CalcularPrecio() float32 {
	return pizza.Precio
}

func NewPizzaSabor(nombre string, precio float32) iPizza {
	return &PizzaSabor{
		Nombre: nombre,
		Precio: precio,
	}
}
