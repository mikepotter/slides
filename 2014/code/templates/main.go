package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Page struct {
	Items map[int]Item
}

type Item struct {
	ID   int
	Name string
}

var (
	temps map[string]*template.Template
	items map[int]Item
)

func init() {
	items = map[int]Item{
		1: Item{1, "Shoes"},
		2: Item{2, "Hats"},
		3: Item{3, "Bacon"},
	}
}

func list(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	p := &Page{items}
	temps["list"].Execute(w, p)
}

func show(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	strid := q.Get("id")
	id, err := strconv.Atoi(strid)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	item, ok := items[id]
	if !ok {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	temps["show"].Execute(w, item)
}

func add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.PostFormValue("name")
		strid := r.PostFormValue("id")
		id, err := strconv.Atoi(strid)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}
		items[id] = Item{id, name}
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	temps["new"].Execute(w, struct{}{})
}

func del(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	strid := q.Get("id")
	id, err := strconv.Atoi(strid)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	delete(items, id)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func main() {
	l, _ := template.ParseFiles("list.html")
	s, _ := template.ParseFiles("show.html")
	n, _ := template.ParseFiles("new.html")
	temps = map[string]*template.Template{
		"list": l,
		"show": s,
		"new":  n,
	}

	http.HandleFunc("/", list)
	http.HandleFunc("/show", show)
	http.HandleFunc("/new", add)
	http.HandleFunc("/delete", del)

	log.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
