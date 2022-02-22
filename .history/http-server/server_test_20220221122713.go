package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGETPlayers(t *testing.T) {

	check := assert.New(t)
	t.Run("returns Pepper's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
		response := httptest.NewRecorder()
	
		PlayerServer(response, request)
	
		got := response.Body.String()
		want := "20"
	
		check.Equal(want, got)
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Floyd", nil)
		response := httptest.NewRecorder()
	
		PlayerServer(response, request)
	
		got := response.Body.String()
		want := "10"
	
		check.Equal(want, got)
	})
}