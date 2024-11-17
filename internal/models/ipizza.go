package models

type IPizza interface {
	CalcularPrecio() float32
	GetTipo() string
}
