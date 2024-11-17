package api

import (
	"net/http"

	"github.com/carlosclavijo/pizza/internal/handlers"
)

func routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /ingredientes", handlers.Repo.PostIngrediente)
	mux.HandleFunc("GET /ingredientes", handlers.Repo.GetIngredientes)
	mux.HandleFunc("GET /ingredientes/{id}", handlers.Repo.GetIngredienteById)

	mux.HandleFunc("POST /sabores", handlers.Repo.PostSabor)
	mux.HandleFunc("GET /sabores", handlers.Repo.GetSabores)
	mux.HandleFunc("GET /sabores/{id}", handlers.Repo.GetSaborById)

	mux.HandleFunc("POST /pedidos", handlers.Repo.PostPedido)

	return mux
}
