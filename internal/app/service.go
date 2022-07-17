package app

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
)

type Service struct {
	Ip     string
	Port   string
	Router *chi.Mux
	DB     *sql.DB
}
