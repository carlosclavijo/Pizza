package handlers

import "gorm.io/gorm"

type Repository struct {
	Db *gorm.DB
}

var Repo *Repository

func NewRepo(db *gorm.DB) *Repository {
	return &Repository{
		Db: db,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}
