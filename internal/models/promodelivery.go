package models

import "reflect"

type PromoDelivery struct {
	Nombre string
}

func (PromoDelivery) Descuento(pedido *Pedido) float32 {
	descuento := pedido.PrecioDelivery
	pedido.PrecioDelivery = 0
	return descuento
}

func (p *PromoDelivery) GetTipo() string {
	return reflect.TypeOf(p).Name()
}

func NewPromoDelivery() IPromocion {
	return &PromoDelivery{
		Nombre: "Delivery Gratis",
	}
}
