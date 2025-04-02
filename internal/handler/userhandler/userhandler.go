package userhandler

import (
	"encoding/json"
	"errors"
	"golangproj/internal/domain/userdomain"
	"golangproj/internal/usecase/user"
	"golangproj/internal/utils"
	"net/http"
	"strconv"
)

type UserHandler struct {
	Usecases *user.UserUsecases
}

func (u *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var input userdomain.UserInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	if err := input.Validate(); err != nil {
		utils.RespondWithError(w, http.StatusUnprocessableEntity, err)
		return
	}

	response, err := u.Usecases.Create.Execute(r.Context(), &input)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, response)
}

func (u *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, errors.New("Error getting id. Check if its not null or empty"))
		return
	}

	response, err := u.Usecases.GetById.Execute(r.Context(), id)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err)
	}

	utils.RespondWithJSON(w, http.StatusOK, response)
}
