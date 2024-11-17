package api

import (
	"log"
	"net/http"

	"gorm.io/gorm"
)

type Api struct {
	port string
	db   *gorm.DB
}

func NewApi(port string, db *gorm.DB) *Api {
	return &Api{
		port: port,
		db:   db,
	}
}

func (s *Api) Run() error {
	server := &http.Server{
		Addr:    s.port,
		Handler: routes(),
	}
	log.Printf("Server has started %s", s.port)
	return server.ListenAndServe()
}
