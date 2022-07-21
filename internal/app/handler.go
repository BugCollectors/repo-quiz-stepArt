package app

import (
	"github.com/go-chi/chi/v5"
	"study/internal/service"
)

//type Handler func(w http.ResponseWriter, r *http.Request) error

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() chi.Router {
	r := chi.NewRouter()

	r.Route("/user", func(r chi.Router) {
		r.Post("/post", h.createUser) // "/user/post"

		//r.Route("/{id}", func(r chi.Router) {
		r.Get("/", h.getUser)
		//})

	})

	return r
}
