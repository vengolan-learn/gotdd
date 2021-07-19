package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) RecordWin(name string) {

}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case http.MethodPost:
		p.ProcessWin(w, r, player)
	case http.MethodGet:
		p.showScore(w, r, player)
	}

}

func (p *PlayerServer) showScore(w http.ResponseWriter, r *http.Request, player string) {

	score := p.store.GetPlayerScore(player)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, score)
}

func (p *PlayerServer) ProcessWin(w http.ResponseWriter, r *http.Request, player string) {

	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

/* func GetPlayerScore(player string) int {
	switch player {
	case "Pepper":
		return 20
	case "Floyd":
		return 10
	default:
		return 0
	}
}
*/
