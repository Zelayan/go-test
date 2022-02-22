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

	switch r.
	shouldReturn := processWin(r, w)
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

func (p *PlayerServer) processWin(w http.ResponseWriter) {

	w.WriteHeader(http.StatusAccepted)

}


type StubPlayerStore struct {
    scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
    score := s.scores[name]

    return score
}