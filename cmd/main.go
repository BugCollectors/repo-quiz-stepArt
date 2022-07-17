package main

import (
	"log"
	"net/http"
)

func main() {
	service := createNewService()
	log.Fatalln(http.ListenAndServe(":3333", service.Router))
}
