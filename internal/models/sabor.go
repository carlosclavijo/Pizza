package models

type Sabor struct {
	SaborId int     `json:"sabor_id" gorm:"unique;primaryKey;autoIncrement"`
	Nombre  string  `json:"nombre"`
	Precio  float32 `json:"precio"`
}
