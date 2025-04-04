package authhandler

import (
	"encoding/json"
	"morae/internal/domain/authdomain"
	"morae/internal/dto/userdto"
	"morae/internal/env"
	"morae/internal/jwt"
	"morae/internal/usecase/auth"
	"morae/internal/utils"
	"net/http"
	"time"
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
	var input userdto.UserInput
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

func (a *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     env.GetString("AUTH_TOKEN_NAME", "dev_token"),
		Value:    "",
		Path:     "/",
		MaxAge:   -1, // tells browser to delete the cookie
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		Secure:   env.GetBool("SECURE_TOKEN", false), // Must match
		SameSite: http.SameSiteStrictMode,
	})

	utils.RespondWithSuccess(w, http.StatusOK, "Logged out successfully")
}
