package main

import (
	"log"
	"net/http"
)

func main() {
	server := NewPl{NewInMemoryPlayerStore()}
	log.Fatal(http.ListenAndServe(":5000", server))
}