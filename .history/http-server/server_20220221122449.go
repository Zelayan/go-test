package main

import (
	"fmt"
	"net/http"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {

	player := r.URL.Path[len("/player/"):]

	if player == "Pepper" {
		fmt.Fprintf(w, "20")
		return
	}

	if player == "Floyd" {
		fmt.Fprintf(w, "10")
		return
	}

	fmt.Fprintxx(w, "20")
}