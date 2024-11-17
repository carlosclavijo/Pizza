package models

type Promo2x1 struct{}

func (Promo2x1) Descuento(pedido *Pedido) {
	pedido.precioPizzas = pedido.precioPizzas / 2
}

func NewPromo2x1(ingredientes []Ingrediente) iPromocion {
	return &Promo2x1{}
}
