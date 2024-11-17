package models

type iPromocion interface {
	Descuento(pedido *Pedido)
}
