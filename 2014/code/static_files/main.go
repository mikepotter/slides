package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("/Users/mike/Desktop")))
	http.HandleFunc("/myfile", serveMyFile)

	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func serveMyFile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/Users/mike/Desktop/go_development_environment_setup.md")
}
