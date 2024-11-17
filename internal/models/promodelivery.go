package models

type PromoDelivery struct{}

func (PromoDelivery) Descuento(pedido *Pedido) {
	pedido.precioDelivery = 0
}

func NewPromoDelivery() iPromocion {
	return &PromoDelivery{}
}
