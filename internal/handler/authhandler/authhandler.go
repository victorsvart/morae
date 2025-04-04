package authhandler

import (
	"encoding/json"
	"morae/internal/domain/authdomain"
	"morae/internal/domain/userdomain"
	"morae/internal/env"
	"morae/internal/jwt"
	"morae/internal/usecase/auth"
	"morae/internal/utils"
	"net/http"
)

type AuthHandler struct {
	Usecases *auth.AuthUsecases
}

func (a *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var input authdomain.LoginInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.RespondWithError(w, http.StatusUnprocessableEntity, err)
		return
	}

	err := a.Usecases.Login.Execute(r.Context(), &input)
	if err != nil {
		utils.RespondWithError(w, http.StatusUnauthorized, err)
		return
	}

	token, err := jwt.GenerateJWT(input.EmailAddress)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     env.GetString("AUTH_TOKEN_NAME", "dev_token"),
		Value:    token,
		HttpOnly: true,
		Secure:   env.GetBool("SECURE_TOKEN", false), // should be true in prod
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
		MaxAge:   3000,
	})

	utils.RespondWithSuccess(w, http.StatusOK, "Logged in successfully")
}

func (a *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var input userdomain.UserInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	err := a.Usecases.Register.Execute(r.Context(), &input)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	utils.RespondWithSuccess(w, http.StatusCreated, "Registered successfully")
}
