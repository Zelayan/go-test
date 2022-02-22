package main

import "net/http"

func main() {
	handle := http.HandlerFunc(PlayerServer)
	if err := http.ListenAndServe(":5000", handle)
}