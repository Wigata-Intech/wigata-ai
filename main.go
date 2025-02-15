package main

import (
	"log"
	"net/http"
	"wigata-ai/handlers"
)

func main() {
	routes := handlers.New()

	log.Println("Server started at :3000")
	http.ListenAndServe(":3000", routes)
}
