package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/tckthecreator/go-api-with-pgsql/models"
)

func List(w http.ResponseWriter, r *http.Request) {
	todos, err := models.GetAll()

	if err != nil {
		log.Printf("Erro ao listar registros %v\n", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
