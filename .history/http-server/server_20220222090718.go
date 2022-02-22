package main

import (
	"fmt"
	"net/http"
)


type PlayerStore interface {
    GetPlayerScore(name string) int
}


type PlayerServer struct {
    store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    player := r.URL.Path[len("/players/"):]
	w.WriteHeader(http.StatusNotFound)

	score := p.store.GetPlayerScore(player)
    fmt.Fprint(w, )
}


type StubPlayerStore struct {
    scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
    score := s.scores[name]

    return score
}