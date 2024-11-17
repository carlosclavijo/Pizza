package models

import "reflect"

type Promo2x1 struct {
	Nombre string
}

func (Promo2x1) Descuento(pedido *Pedido) float32 {
	descuento := pedido.PrecioPizzas / 2
	return descuento
}

func (p *Promo2x1) GetTipo() string {
	return reflect.TypeOf(p).Name()
}

func NewPromo2x1() IPromocion {
	return &Promo2x1{
		Nombre: "Promocion 2x1",
	}
}
