package main

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"study/internal/app"
)

type Handler func(w http.ResponseWriter, r *http.Request) error

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		// handle returned error here.
		w.WriteHeader(503)
		w.Write([]byte("bad"))
	}
}

func createNewService() *app.Service {
	s := &app.Service{}

	s.Router = chi.NewRouter()
	s.Router.Method("GET", "/", Handler(s.Get))
	s.Router.Method("POST", "/post", Handler(s.Post))
	s.Router.Method("PUT", "/put", Handler(s.Put))
	s.Router.Method("DELETE", "/delete", Handler(s.Delete))

	//s.DB = createDbConnect()

	return s
}

func createDbConnect() *sql.DB {
	connStr := "postgresql://<username>:<password>@<database_ip>/todos?sslmode=disable"
	// Connect to database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
