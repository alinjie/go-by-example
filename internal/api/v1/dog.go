package v1

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type DogDTO struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type DogStorage interface {
	ListDogs() ([]DogDTO, error)
	GetDog(id int) (DogDTO, error)
}

type DogHandler struct {
	storage DogStorage
}

func NewDogHandler(s DogStorage) DogHandler {
	return DogHandler{
		storage: s,
	}
}

func (h DogHandler) HandleListDogs(w http.ResponseWriter, r *http.Request) {
	dogs, err := h.storage.ListDogs()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errInternalServerError))
		return
	}

	j, err := json.Marshal(dogs)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errInternalServerError))
		return
	}

	w.Write(j)
}

func (h DogHandler) HandleGetDog(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errInvalidIdParam))
		return
	}

	dog, err := h.storage.GetDog(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errInternalServerError))
		return
	}

	j, err := json.Marshal(dog)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errInternalServerError))
		return
	}

	w.Write(j)
}
