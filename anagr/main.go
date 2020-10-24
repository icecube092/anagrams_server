package main

import (
	"anagrams_server/anagr/handlers"
	"log"
	"net/http"
)

var addr = "localhost:8080"

func main() {
	server := http.Server{}
	server.Addr = addr
	handler := &handlers.HttpHandler{}
	handler.Storage = make(map[string][]string)
	server.Handler = handler

	defer func() {
		log.Println("Exit...")
		err := server.Close()
		if err != nil {
			log.Println(err)
		}
	}()
	log.Println("Listening...")
	err := server.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}
