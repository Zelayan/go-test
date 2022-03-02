package main

import (
	"log"
	"net/http"
)

func main() {
	server := NewPlayerServer{&st}
	log.Fatal(http.ListenAndServe(":5000", server))
}
