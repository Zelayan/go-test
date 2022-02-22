package main

import (
	"log"
	"net/http"
)

func main() {
	handle := http.HandlerFunc(PlayerServer)
	if err := http.ListenAndServe(":5000", handle); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}