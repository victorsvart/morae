package userhandler

import (
	"encoding/json"
	"errors"
	"morae/internal/dto/userdto"
	"morae/internal/mapper/usermapper"
	"morae/internal/usecase/user"
	"morae/internal/utils"
	"net/http"
	"strconv"
)

type UserHandler struct {
	Usecases *user.UserUsecases
}

func (u *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest,
			errors.New("Error getting id. Check if its not null or empty"))

		return
	}

	if id == 0 {
		utils.RespondWithError(w, http.StatusBadRequest, ErrInvalidId)
	}

	response, err := u.Usecases.GetById.Execute(r.Context(), id)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, err)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, response)
}

func (u *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var input userdto.UserInput

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

func (u *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var userDto userdto.UserDto

	if err := json.NewDecoder(r.Body).Decode(&userDto); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	domain, err := usermapper.FromDto(&userDto)
	if err != nil {
		utils.RespondWithError(w, http.StatusUnprocessableEntity, err)
		return
	}

	if domain.Password.Value != "" {
		if err := domain.Password.HashPassword(); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}
	}
	response, err := u.Usecases.Update.Execute(r.Context(), domain)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, response)
}

func (u *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, errors.New("invalid user ID"))
		return
	}

	err = u.Usecases.Delete.Execute(r.Context(), id)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	utils.RespondWithSuccess(w, http.StatusOK, "User deleted successfully")
}
