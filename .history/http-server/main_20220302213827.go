package main

import (
	"log"
	"net/http"
)

func main() {
	imps := InMemoryPlayerStore{}
	server := NewPlayerServer(imps)
	log.Fatal(http.ListenAndServe(":5000", server))
}
