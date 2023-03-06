package server

import (
	"net/http"
	"wwchacalww/go-psyc/adapters/http/handler"
	"wwchacalww/go-psyc/domain/repository"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
)

type WebServer struct {
	AuthRepository repository.AuthRepositoryInterface
	UserRepository repository.UserRepositoryInterface
}

func MakeNewWebServer() *WebServer {
	return &WebServer{}
}

func (w WebServer) Server() {
	r := chi.NewRouter()
	c := cors.New(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(c.Handler)
	r.Use(middleware.Logger)
	handler.MakeAuthHandlers(r, w.AuthRepository)
	handler.MakeUserHandlers(r, w.UserRepository)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	http.ListenAndServe(":9000", r)
}
