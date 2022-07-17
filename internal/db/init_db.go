package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"log"
)

func CreateDB() *sqlx.DB {

	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		fmt.Sprint(viper.Get("POSTGRES_HOST")),
		fmt.Sprint(viper.Get("POSTGRES_PORT")),
		fmt.Sprint(viper.Get("POSTGRES_USERNAME")),
		fmt.Sprint(viper.Get("POSTGRES_PASSWORD")),
		fmt.Sprint(viper.Get("POSTGRES_DBNAME")),
		fmt.Sprint(viper.Get("POSTGRES_SSLMODE")),
	)

	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
