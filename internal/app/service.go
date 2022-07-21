package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v4"
)

type Service struct {
	Ip     string
	Port   string
	Router *chi.Mux
	DB     *pgx.Conn
}
