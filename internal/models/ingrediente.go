package models

type Ingrediente struct {
	IngredienteId int     `json:"ingrediente_id" gorm:"unique;primaryKey;autoIncrement"`
	Nombre        string  `json:"nombre"`
	Precio        float32 `json:"precio"`
}
