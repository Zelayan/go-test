package main

import (
	"fmt"
	"net/http"
)

type PlayerServer struct {
    store PlayerStore
}


type PlayerStore interface {
    GetPlayerScore(name string) int
}


func ()PlayerServer(w http.ResponseWriter, r *http.Request) {

	player := r.URL.Path[len("/players/"):]

	if player == "Pepper" {
		fmt.Fprintf(w, "20")
		return
	}

	if player == "Floyd" {
		fmt.Fprintf(w, "10")
		return
	}
}