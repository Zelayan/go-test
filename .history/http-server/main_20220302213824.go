package main

import (
	"log"
	"net/http"
)

func main() {
	imps := InMemoryPlayerStore{}
	server := NewPlayerServer()
	log.Fatal(http.ListenAndServe(":5000", server))
}
