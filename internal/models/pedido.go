package models

import (
	"time"
)

// Estado 1 = En proceso, 2 = Pagado, 3 = Cancelado
type Pedido struct {
	PedidoId       int `json:"pedido_id" gorm:"unique;primaryKey;autoIncrement"`
	Fecha          time.Time
	Pizzas         []IPizza   `json:"pizza" gorm:"-"`
	Promocion      IPromocion `json:"promociones" gorm:"-"`
	PrecioPizzas   float32
	PrecioDelivery float32
	Descuento      float32
	PrecioTotal    float32
}
