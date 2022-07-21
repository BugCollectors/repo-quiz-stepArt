package main

import (
	"context"
	"github.com/go-chi/chi/v5"
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

func createNewService(ctx context.Context) (*app.Service, error) {
	s := &app.Service{}
	s.App = &app.Application{}

	s.Router = chi.NewRouter()
	s.Router.Method("GET", "/get_chepush_message", Handler(s.App.GetChepushMessageByID))
	s.Router.Method("POST", "/new_message_from_chepush", Handler(s.App.NewMessageFromChepush))
	s.Router.Method("DELETE", "/delete_chepush_message", Handler(s.App.DeleteChepushMessageByID))

	config, err := cfg.LoadConfig()
	if err != nil {
		return nil, err
	}

	postgresConnection, err := db.CreateDB(ctx, config)

	s.App.ChepushilaRepository = db.NewChepushilaRepo(postgresConnection)
	return s, nil
}
