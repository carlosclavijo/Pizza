package models

import "github.com/gofrs/uuid"

type Ingrediente struct {
	IngrdienteId uuid.UUID
	Nombre       string
	Precio       float32
}
