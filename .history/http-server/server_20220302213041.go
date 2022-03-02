package main

import (
	"fmt"
	"net/http"
)

// PlayerStore stores score information about players.
type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

// PlayerServer is a HTTP interface for player information.
type PlayerServer struct {
	store PlayerStore
}

func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
})