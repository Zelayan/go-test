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
	shouldReturn := newFunction(r, w)
	if shouldReturn {
		return
	}

    player := r.URL.Path[len("/players/"):]


	score := p.store.GetPlayerScore(player)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
    fmt.Fprint(w, score)
}

func newFunction(r *http.Request, w http.ResponseWriter) bool {
	if r.Method == http.MethodPost {
		w.WriteHeader(http.StatusAccepted)
		return true
	}
	return false
}


type StubPlayerStore struct {
    scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
    score := s.scores[name]

    return score
}