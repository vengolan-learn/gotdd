package main

import (
	"log"
	"net/http"
)

func main() {
	server := NewPlayerServer(NewInMemoryPlayerStore())
	// handler := http.HandlerFunc(PlayerServer)
	log.Fatal(http.ListenAndServe(":5000", server))
}
