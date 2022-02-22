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

	switch r.Method {
	case http.MethodGet:
		p.showScore(w, r)
	case http.MethodPost:
		p.processWin(w)
	}
}

func (p *PlayerServer) showScore(w http.ResponseWriter, r *http.Request) {
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
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
    score := s.scores[name]

    return score
}

func (s *StubPlayerStore) RecordWin(name string) {
    s.winCalls = append(s.winCalls, name)
}