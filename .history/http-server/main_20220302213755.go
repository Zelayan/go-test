package main

import (
	"log"
	"net/http"
)

func main() {
	InMemoryPlayerStore
	server := NewPlayerServer(InMemoryPlayerStore)
	log.Fatal(http.ListenAndServe(":5000", server))
}
