package main

import (
	"io"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		dump, _ := httputil.DumpRequest(r, true)
		io.WriteString(w, string(dump))
	})

	log.Println("Listening on localhost: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
