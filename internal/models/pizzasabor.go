package models

import "reflect"

type PizzaSabor struct {
	PizzaSaborId int   `json:"pizzareceta_id" gorm:"unique;primaryKey;autoIncrement"`
	SaborId      int   `json:"sabor_id"`
	PedidoId     int   `json:"pedido_id"`
	Sabor        Sabor `gorm:"-"`
}

func (p PizzaSabor) CalcularPrecio() float32 {
	return p.Sabor.Precio
}

func (p PizzaSabor) GetTipo() string {
	return reflect.TypeOf(p).Name()
}

func NewPizzaSabor(sabor, pedido int) IPizza {
	return &PizzaSabor{
		SaborId:  sabor,
		PedidoId: pedido,
	}
}
