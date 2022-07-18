package app

import (
	"github.com/go-chi/chi/v5"
	"study/internal/service"
)

//type Handler func(w http.ResponseWriter, r *http.Request) error

type Router struct {
	Router   *chi.Mux
	services *service.Service
}

func NewHandler(services *service.Service) *Router {
	return &Router{services: services}
}

func (h *Router) InitRoutes() {
	h.Router = chi.NewRouter()
	h.Router.Post("/post", h.services.CreateUser)
}
