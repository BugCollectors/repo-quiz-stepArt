package main

import (
	"log"
	"net/http"
)

//1. Сделать конфиг
//2. Подключение к базе
//3. Метод записи

func main() {
	service, err := createNewService()
	if err != nil {
		log.Fatalln(err)
	}
	log.Fatalln(http.ListenAndServe(":3333", service.Router))
}
