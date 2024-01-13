package userController

import (
	userModel "api-crud/internal/entities/user/model"
	errorModel "api-crud/internal/shared/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userId := chi.URLParam(r, "id");
	userName := chi.URLParam(r, "name");
	
	id, err := strconv.Atoi(userId)
	
	if err != nil {
		errorMessage := errorModel.AppError {
			Message: "Id inválido",
		}
		error, _ := json.Marshal(errorMessage)
		http.Error(w, string(error), http.StatusBadRequest)
		return
	}

	user := userModel.User {
		ID: id,
		Name: userName,
	}

	data, err := json.Marshal(user)

	if err != nil {
		errorMessage := errorModel.AppError {
			Message: "Erro ao gerar JSON do Usuário.",
		}
		error, _ := json.Marshal(errorMessage)
		http.Error(w, string(error), http.StatusInternalServerError)
		return
	}
	
	w.Write(data);
}