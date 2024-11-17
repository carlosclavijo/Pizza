package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Pedido struct {
	pedidoId       uuid.UUID
	fecha          time.Time
	pizzas         []iPizza
	promocion      iPromocion
	precioPizzas   float32
	precioDelivery float32
	precioTotal    float32
}
