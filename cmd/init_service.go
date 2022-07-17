package main

import (
	"log"
	"net/http"
	"os"

	"database/sql"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"study/internal/app"
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

func createNewService() *app.Service {
	s := &app.Service{}

	s.Router = chi.NewRouter()
	s.Router.Method("GET", "/", Handler(s.Get))
	s.Router.Method("POST", "/post", Handler(s.Post))
	s.Router.Method("PUT", "/put", Handler(s.Put))
	s.Router.Method("DELETE", "/delete", Handler(s.Delete))

	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error init config: %v", err)
	}

	s.DB, _ = db.NewPostgresDB(db.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD")},
	)

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

func initConfig() error {
	viper.AddConfigPath("cfg")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
