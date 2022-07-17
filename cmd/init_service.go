package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"
	"net/http"
	"study/internal/app"
	"study/internal/cfg"
	"study/internal/db"
)

type Handler func(w http.ResponseWriter, r *http.Request) error

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		// handle returned error here.
		w.WriteHeader(503)
		w.Write([]byte("bad"))
	}
}

func createNewService() (*app.Service, error) {
	s := &app.Service{}

	s.Router = chi.NewRouter()
	s.Router.Method("GET", "/", Handler(s.Get))
	s.Router.Method("POST", "/post", Handler(s.Post))
	s.Router.Method("PUT", "/put", Handler(s.Put))
	s.Router.Method("DELETE", "/delete", Handler(s.Delete))

	err := cfg.ParseConfig()
	if err != nil {
		return nil, err
	}

	s.DB = db.CreateDB()

	return s, nil
}
