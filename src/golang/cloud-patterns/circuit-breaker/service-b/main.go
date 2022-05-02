package main

import (
	"log"
	"service-b/handlers"
)

func main() {
	handler := handlers.NewHandler()
	go log.Fatal(handler.ListenAndServe())
}
