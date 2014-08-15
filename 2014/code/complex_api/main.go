package main

import (
	"log"
	"net/http"
	"time"

	"./client"
)

type MainHandler struct {
	client http.Handler
}

func (h *MainHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var handler http.Handler
	switch r.URL.Path {
	case "/api/clients":
		handler = h.client
	default:
		http.NotFound(w, r)
		return
	}
	handler.ServeHTTP(w, r)
}

func main() {
	mainHandler := &MainHandler{
		client: client.Handler(),
	}

	s := &http.Server{
		Addr:           ":8080",
		Handler:        mainHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Listening on port 8080")
	log.Fatal(s.ListenAndServe())
}
