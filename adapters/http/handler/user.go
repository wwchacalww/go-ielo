package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"wwchacalww/go-psyc/domain/repository"
	"wwchacalww/go-psyc/domain/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

type UserHandler struct {
	UserRepository repository.UserRepositoryInterface
}

func MakeUserHandlers(r *chi.Mux, repo repository.UserRepositoryInterface) {
	handler := &UserHandler{
		UserRepository: repo,
	}
	jwtoken := jwtauth.New("HS256", []byte("secret_jwt"), nil)

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(jwtoken))
		r.Use(jwtauth.Authenticator)
		r.Route("/users", func(r chi.Router) {
			r.Post("/", handler.Store)
			r.Get("/{id}", handler.FindById)
			r.Get("/email", handler.FindByEmail)
			r.Get("/", handler.List)
			r.Put("/change/password", handler.ChangePassword)
			r.Put("/change/my/password", handler.ChangeMyPassword)
			r.Put("/change/role", handler.ChangeRole)
			r.Put("/change/avatar", handler.ChangeAvatarUrl)
			r.Put("/change/my/avatar", handler.ChangeMyAvatarUrl)
		})
	})

}

func (u *UserHandler) Store(w http.ResponseWriter, r *http.Request) {
	token, _, _ := jwtauth.FromContext(r.Context())
	role, _ := token.Get("role")
	if role != "Admin" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(jsonError("Access denied"))
		return
	}
	var input repository.UserInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}

	user, err := u.UserRepository.Create(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (u *UserHandler) FindById(w http.ResponseWriter, r *http.Request) {
	token, _, _ := jwtauth.FromContext(r.Context())
	role, _ := token.Get("role")
	if role != "Admin" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(jsonError("Access denied"))
		return
	}
	id := chi.URLParam(r, "id")
	user, err := u.UserRepository.FindById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (u *UserHandler) FindByEmail(w http.ResponseWriter, r *http.Request) {
	token, _, _ := jwtauth.FromContext(r.Context())
	tokenEmail, _ := token.Get("email")
	email := fmt.Sprintf("%v", tokenEmail)

	user, err := u.UserRepository.FindByEmail(email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (u *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	token, _, _ := jwtauth.FromContext(r.Context())
	role, _ := token.Get("role")
	if role != "Admin" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(jsonError("Access denied"))
		return
	}
	users, err := u.UserRepository.List()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (u *UserHandler) ChangePassword(w http.ResponseWriter, r *http.Request) {
	token, _, _ := jwtauth.FromContext(r.Context())
	role, _ := token.Get("role")
	if role != "Admin" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(jsonError("Access denied"))
		return
	}
	var input struct {
		ID       string `json:"id"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}

	err = u.UserRepository.ChangePassword(input.ID, input.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}

	w.WriteHeader(201)
	w.Write(jsonError("Password updated"))
}

func (u *UserHandler) ChangeMyPassword(w http.ResponseWriter, r *http.Request) {
	token, _, _ := jwtauth.FromContext(r.Context())
	tokenEmail, _ := token.Get("email")
	email := fmt.Sprintf("%v", tokenEmail)

	user, err := u.UserRepository.FindByEmail(email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}

	var input struct {
		Password string `json:"password"`
	}
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}

	err = u.UserRepository.ChangePassword(user.GetID(), input.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}

	w.WriteHeader(201)
	w.Write(jsonError("Password updated"))
}

func (u *UserHandler) ChangeRole(w http.ResponseWriter, r *http.Request) {
	token, _, _ := jwtauth.FromContext(r.Context())
	role, _ := token.Get("role")
	if role != "Admin" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(jsonError("Access denied"))
		return
	}
	var input struct {
		ID   string `json:"id"`
		Role string `json:"role"`
	}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}

	err = u.UserRepository.ChangeRole(input.ID, input.Role)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}

	w.WriteHeader(201)
	w.Write(jsonError("Role updated"))
}

func (u *UserHandler) ChangeMyAvatarUrl(w http.ResponseWriter, r *http.Request) {
	token, _, _ := jwtauth.FromContext(r.Context())
	tokenEmail, _ := token.Get("email")
	tokenAvatar, _ := token.Get("avatar_url")
	email := fmt.Sprintf("%v", tokenEmail)
	oldAvatar := fmt.Sprintf("%v", tokenAvatar)
	log.Println(oldAvatar)
	os.Remove("public/imgs/" + oldAvatar)
	f, fh, err := r.FormFile("avatar")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}

	err = utils.AvatarIsValid(fh)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}

	defer f.Close()
	avatar_url, err := u.UserRepository.ChangeAvatarUrl(email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}

	dst, err := os.Create("./public/imgs/" + avatar_url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	_, err = io.Copy(dst, f)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}

	err = utils.AvatarResize(avatar_url)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}

	w.WriteHeader(201)
	w.Write(jsonError("Avatar updated"))
}

func (u *UserHandler) ChangeAvatarUrl(w http.ResponseWriter, r *http.Request) {
	token, _, _ := jwtauth.FromContext(r.Context())
	role, _ := token.Get("role")
	if role != "Admin" {
		log.Println(role)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(jsonError("Access denied"))
		return
	}

	email := r.FormValue("email")
	user, err := u.UserRepository.FindByEmail(email)
	if err != nil {
		log.Println(email)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}

	os.Remove("public/imgs/" + user.GetAvatarUrl())
	f, fh, err := r.FormFile("avatar")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}

	err = utils.AvatarIsValid(fh)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}

	defer f.Close()
	avatar_url, err := u.UserRepository.ChangeAvatarUrl(email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}

	dst, err := os.Create("./public/imgs/" + avatar_url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	_, err = io.Copy(dst, f)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}

	err = utils.AvatarResize(avatar_url)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError(err.Error()))
		return
	}

	w.WriteHeader(201)
	w.Write(jsonError("Avatar updated"))
}
