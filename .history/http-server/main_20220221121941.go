package main

import "net/http"

func main() {
	http.HandleFunc(PlayerServer())
}