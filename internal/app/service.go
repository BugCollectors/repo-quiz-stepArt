package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

type Service struct {
	Ip     string
	Port   string
	Router *chi.Mux
	DB     *sqlx.DB
}
