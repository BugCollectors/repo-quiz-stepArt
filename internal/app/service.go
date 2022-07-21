package app

import (
	"github.com/go-chi/chi/v5"
	"study/internal/db"
)

type Service struct {
	Ip     string
	Port   string
	Router *chi.Mux
	App    *Application
}

type Application struct {
	ChepushilaRepository db.Chepushila
}
