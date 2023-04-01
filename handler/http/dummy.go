package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type dummyHandler struct {
}

func NewDummyHandler(router *chi.Mux) {
	handler := dummyHandler{}

	router.Get("/", handler.HelloWorld)
}

func (d *dummyHandler) HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte(`{"message": "OK"}`))
}
