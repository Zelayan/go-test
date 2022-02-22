package main

import "net/http"

func main() {
	handle := http.HandlerFunc(PlayerServer)
}