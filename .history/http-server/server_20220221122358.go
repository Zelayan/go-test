package main

import (
	"fmt"
	"net/http"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {

	player := r.URL.Path[len("/player/"):]

	if 

	fmt.Fprint(w, "20")
}