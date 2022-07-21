package main

import (
	"context"
	"log"
	"net/http"
)

//1. Структура простой таблицы и миграция
//2. Подключение к базе
//3. Метод записи

func main() {
	ctx := context.Background()

	service, err := createNewService(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	log.Fatalln(http.ListenAndServe(":3333", service.Router))
}
