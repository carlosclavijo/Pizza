package models

import "reflect"

type PizzaPersonalizada struct {
	PizzaPersonalizadaId int           `json:"pizzapersonalizada_id" gorm:"unique;primaryKey;autoIncrement"`
	Ingredientes         []Ingrediente `gorm:"-"`
	Precio               float32       `json:"precio"`
	PedidoId             int           `json:"pedido_id"`
}

func (p PizzaPersonalizada) CalcularPrecio() float32 {
	p.Precio = 30
	for _, ingrediente := range p.Ingredientes {
		p.Precio += ingrediente.Precio
	}
	return p.Precio
}

func (p *PizzaPersonalizada) GetTipo() string {
	return reflect.TypeOf(p).Name()
}

func NewPizzaPersonalizada(ingredientes []Ingrediente, pedido int) IPizza {
	return &PizzaPersonalizada{
		Ingredientes: ingredientes,
		PedidoId:     pedido,
	}
}
