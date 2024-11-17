package models

type IPromocion interface {
	Descuento(pedido *Pedido) float32
	GetTipo() string
}
