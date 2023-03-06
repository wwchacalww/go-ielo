package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"wwchacalww/go-psyc/domain/repository"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

type AuthHandler struct {
	AuthRepository repository.AuthRepositoryInterface
}

func MakeAuthHandlers(r *chi.Mux, repo repository.AuthRepositoryInterface) {
	handler := &AuthHandler{
		AuthRepository: repo,
	}
	jwtoken := jwtauth.New("HS256", []byte("secret_jwt"), nil)
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(jwtoken))
		r.Post("/auth/refresh_token", handler.RefreshToken)
	})
	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", handler.Authenticate)
	})
}

func (a *AuthHandler) Authenticate(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}

	result, err := a.AuthRepository.Authenticate(input.Email, input.Password)
	if err != nil {
		w.WriteHeader(403)
		w.Write(jsonError(err.Error()))
		return
	}
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (a *AuthHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	token, _, err := jwtauth.FromContext(r.Context())
	log.Println(err.Error())
	if err != nil && err.Error() != "token is expired" {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	tokenEmail, _ := token.Get("email")
	email := fmt.Sprintf("%v", tokenEmail)
	var input struct {
		RefreshToken string `json:"refresh_token"`
	}
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}

	result, err := a.AuthRepository.RefreshToken(email, input.RefreshToken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}
