package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

var (
	addr = flag.String("addr", ":8081", "The address to listen on")
	pass = flag.String("pass", "http://127.0.0.1:8080", "The server to pass through to")
)

type Handler struct {
	proxy *httputil.ReverseProxy
}

func main() {
	srvr, err := url.Parse(*pass)
	if err != nil {
		panic(err)
	}

	handler := &Handler{
		proxy: httputil.NewSingleHostReverseProxy(srvr),
	}

	s := &http.Server{
		Addr:           *addr,
		Handler:        handler,
		ReadTimeout:    1 * time.Second,
		WriteTimeout:   1 * time.Second,
		MaxHeaderBytes: 100,
	}
	log.Printf("Listening on %s and passing to %s\n", *addr, *pass)
	log.Fatal(s.ListenAndServe())

}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Do some mangling!!!!
	r.Header.Set("X-Test", "Ahhh Yea!")
	h.proxy.ServeHTTP(w, r)
}
