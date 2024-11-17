package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/carlosclavijo/pizza/internal/models"
	"github.com/carlosclavijo/pizza/internal/repository"
)

func (repo Repository) PostSabor(w http.ResponseWriter, r *http.Request) {
	var sabor models.Sabor
	err := json.NewDecoder(r.Body).Decode(&sabor)
	if err != nil {
		log.Println("Error decodificando el request", err)
	}
	result := repository.InsertSabor(&sabor)
	w.Header().Set("Content-Type", "application/json")
	if result != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(result.Error())
		return
	}
	json.NewEncoder(w).Encode(sabor)
}

func (repo Repository) GetSaborById(w http.ResponseWriter, r *http.Request) {
	var sabor models.Sabor
	id, _ := strconv.Atoi(r.PathValue("id"))
	result := repository.GetSaborById(&sabor, id)
	w.Header().Set("Content-Type", "application/json")
	if result != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(result.Error())
		return
	}
	json.NewEncoder(w).Encode(sabor)
}

func (repo Repository) GetSabores(w http.ResponseWriter, r *http.Request) {
	var sabores []models.Sabor
	result := repository.GetSabores(&sabores)
	w.Header().Set("Content-Type", "application/json")
	if result != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(result.Error())
		return
	}
	json.NewEncoder(w).Encode(sabores)
}
