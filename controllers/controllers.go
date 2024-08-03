package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go-rest-api/database"
	"go-rest-api/models"
	"net/http"
	"strconv"
)

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	var personalidades []models.Personalidade
	database.DB.Find(&personalidades)
	json.NewEncoder(w).Encode(personalidades)
}

func BuscarPersonalidadePorId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Printf("Id inválido")
		return
	}

	var personalidade models.Personalidade
	database.DB.First(&personalidade, models.Personalidade{Id: id})

	json.NewEncoder(w).Encode(personalidade)
}

func CriarPersonalidade(w http.ResponseWriter, r *http.Request) {
	var personalidade models.Personalidade
	err := json.NewDecoder(r.Body).Decode(&personalidade)
	if err != nil {
		fmt.Printf("Corpo de requisição inválido")
		return
	}

	database.DB.Create(&personalidade)

	json.NewEncoder(w).Encode(personalidade)
}

func DeletarPersonalidadePorId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Printf("Id inválido")
		return
	}

	var personalidade models.Personalidade
	database.DB.Delete(&personalidade, models.Personalidade{Id: id})
	json.NewEncoder(w).Encode(personalidade)
}

func EditarPersonalidadePorId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Printf("Id inválido")
		return
	}
	var personalidade models.Personalidade

	database.DB.First(&personalidade, models.Personalidade{Id: id})
	err = json.NewDecoder(r.Body).Decode(&personalidade)
	if err != nil {
		fmt.Printf("Corpo de requisição inválido")
		return
	}
	database.DB.Save(&personalidade)
	json.NewEncoder(w).Encode(personalidade)
}
