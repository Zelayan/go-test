package main

import (
	"log"
	"net/http"
)

func main() {
	server := New{NewInMemoryPlayerStore()}
	log.Fatal(http.ListenAndServe(":5000", server))
}
