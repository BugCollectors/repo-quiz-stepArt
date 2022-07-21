package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/http"
	handler2 "study/internal/app"
	"study/internal/repos"
	service2 "study/internal/service"
)

func main() {
	if err := initConfig(); err != nil {
		fmt.Printf("error init config: %v\n", err)
	}

	db, err := repos.NewPostgresDB(repos.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: viper.GetString("db.password")},
	)
	if err != nil {
		fmt.Printf("error connect db: %v\n", err)
	}

	repo := repos.NewRepository(db)
	service := service2.NewService(repo)
	handler := handler2.NewHandler(service)

	router := handler.InitRoutes()

	log.Fatalln(http.ListenAndServe(":3333", router))
}

func initConfig() error {
	viper.AddConfigPath(".cfg")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
