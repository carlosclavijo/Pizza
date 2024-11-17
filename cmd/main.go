package main

import (
	"log"

	"github.com/carlosclavijo/pizza/cmd/api"
	"github.com/carlosclavijo/pizza/internal/database"
	"github.com/carlosclavijo/pizza/internal/handlers"
)

const portNumber = ":8080"

func main() {
	db := database.GetInstance()
	repo := handlers.NewRepo(db.Db)
	handlers.NewHandlers(repo)
	server := api.NewApi(portNumber, db.Db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
