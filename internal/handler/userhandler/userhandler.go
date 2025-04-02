package userhandler

import (
	"encoding/json"
	"golangproj/internal/domain/userdomain"
	"golangproj/internal/usecase/user"
	"net/http"
)

type UserHandler struct {
	Usecases *user.UserUsecases
}

func (u *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var input userdomain.UserInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	response, err := u.Usecases.Create.Execute(r.Context(), &input)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
