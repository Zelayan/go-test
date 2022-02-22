package main

import (
	"fmt"
	"net/http"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {

	r.URL.Path[len("/player/"):]
	fmt.Fprint(w, "20")
}