package v1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Storage interface {
	DogStorage
}

const (
	errInternalServerError = `{"error": "Internal server error"}`
)

func Server(s Storage) chi.Router {
	dogHandler := NewDogHandler(s)

	r := chi.NewRouter()
	r.Use(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "application/json")
			h.ServeHTTP(w, r)
		})
	})
	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/dogs", dogHandler.HandleListDogs)
		r.Get("/dogs/{id}", dogHandler.HandleGetDog)
	})

	return r
}
