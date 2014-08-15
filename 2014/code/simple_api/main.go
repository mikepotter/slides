package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var f func(http.ResponseWriter, *http.Request)
	switch r.Method {
	case "GET":
		f = list
	case "POST":
		f = add
	default:
		f = http.NotFound
	}
	f(w, r)
}

func list(w http.ResponseWriter, r *http.Request) {
	clients, err := allclients() // Get a list of clients
	if err != nil {
		http.Error(w, "Unknown Error", http.StatusInternalServerError)
	}
	w.Write(clients)
}

func add(w http.ResponseWriter, r *http.Request) {
	// Add your client here, and redirect to client list.
	http.Redirect(w, r, "/api/clients", http.StatusSeeOther)
}

func main() {
	http.HandleFunc("/api/clients", ServeHTTP)
	http.Handle("/", http.FileServer(http.Dir("/Users/mike/Desktop")))
	http.HandleFunc("/myfile", serveMyFile)

	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func serveMyFile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/Users/mike/Desktop/go_development_environment_setup.md")
}

// All my hidden stuff to manage clients
type client struct {
	Name string
	ID   int
}

func allclients() ([]byte, error) {

	c1 := client{"My First Client", 1}
	c2 := client{"Another Great Client", 2}
	c3 := client{"Final Client", 3}

	cl := []client{c1, c2, c3}
	return json.Marshal(cl)
}
