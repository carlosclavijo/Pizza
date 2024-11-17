package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/carlosclavijo/pizza/internal/models"
	"github.com/carlosclavijo/pizza/internal/repository"
)

func (repo Repository) PostIngrediente(w http.ResponseWriter, r *http.Request) {
	var ingrediente models.Ingrediente
	err := json.NewDecoder(r.Body).Decode(&ingrediente)
	if err != nil {
		log.Println("Error decodificando el request", err)
	}
	result := repository.InsertIngrediente(&ingrediente)
	w.Header().Set("Content-Type", "application/json")
	if result != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(result.Error())
		return
	}
	json.NewEncoder(w).Encode(ingrediente)
}

func (repo Repository) GetIngredienteById(w http.ResponseWriter, r *http.Request) {
	var ingrediente models.Ingrediente
	id, _ := strconv.Atoi(r.PathValue("id"))
	result := repository.GetIngredientById(&ingrediente, id)
	w.Header().Set("Content-Type", "application/json")
	if result != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(result.Error())
		return
	}
	json.NewEncoder(w).Encode(ingrediente)
}

func (repo Repository) GetIngredientes(w http.ResponseWriter, r *http.Request) {
	var ingredientes []models.Ingrediente
	result := repository.GetIngredientes(&ingredientes)
	w.Header().Set("Content-Type", "application/json")
	if result != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(result.Error())
		return
	}
	json.NewEncoder(w).Encode(ingredientes)
}
