package main

import (
	"log"
	"net/http"

	"github.com/coreos/etcd/store"
)

func main() {
	server := NewPlayerServer{&store}
	log.Fatal(http.ListenAndServe(":5000", server))
}
