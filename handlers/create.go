package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/tckthecreator/go-api-with-pgsql/models"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)

	if err != nil {
		log.Printf("Erro ao fazer decode do JSON %v\n", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := models.Insert(todo)

	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"Error":      true,
			"StatusCode": http.StatusInternalServerError,
			"Message":    fmt.Sprintf("Ocorreu um erro ao tentar inserir: %v\n", err),
		}
	} else {
		resp = map[string]any{
			"Error":      false,
			"StatusCode": http.StatusCreated,
			"Message":    fmt.Sprintf("Todo inserido com sucesso! ID: %d", id),
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
