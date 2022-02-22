package main

import "net/http"

func main() {
	handle := http.HandleFunc(PlayerServer)
}