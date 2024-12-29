package main

import (
	"fmt"
	"log"
	"net/http"
)

const WebPort = "80"

type Config struct{}

func main() {
	app := Config{}

	log.Println("Starting server on port", WebPort)
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", WebPort),
		Handler: app.Routes(),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
		return
	}
}
