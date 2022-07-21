package app_layer

import (
	"github.com/go-chi/chi/v5"
	"study/internal/db_layer"
)

type Service struct {
	Ip     string
	Port   string
	Router *chi.Mux
	App    *Application
}

type Application struct {
	ChepushilaRepository db_layer.Chepushila
}
