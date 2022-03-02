package main

import (
	"log"
	"net/http"
)

func main() {
	server := {NewInMemoryPlayerStore()}
	log.Fatal(http.ListenAndServe(":5000", server))
}