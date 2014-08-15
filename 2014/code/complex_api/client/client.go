package client

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/mediaFORGE/goserve/path"
)

var ps = path.NewScanner("/", "/api/clients/%s")

type Client struct {
	clients map[string]client
	mu      sync.Mutex
}

type client struct {
	Name string
	ID   string
}

func Handler() *Client {
	return &Client{
		clients: map[string]client{},
	}
}

func (c *Client) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var f func(http.ResponseWriter, *http.Request)
	switch r.Method {
	case "GET":
		f = c.list
	case "POST":
		f = c.add
	case "DELETE":
		f = c.del
	default:
		f = http.NotFound
	}
	f(w, r)
}

func (c *Client) list(w http.ResponseWriter, r *http.Request) {
	cl := []client{}
	for _, c := range c.clients {
		cl = append(cl, c)
	}
	list, err := json.Marshal(cl)
	if err != nil {
		http.Error(w, "Unknown Error", http.StatusInternalServerError)
	}
	w.Write(list)
}

func (c *Client) add(w http.ResponseWriter, r *http.Request) {
	name := r.PostFormValue("name")
	id := r.PostFormValue("id")
	if name == "" || id == "" {
		http.Error(w, "Name or ID blank", http.StatusInternalServerError)
		return
	}
	c.mu.Lock()
	c.clients[id] = client{name, id}
	c.mu.Unlock()
	http.Redirect(w, r, "/api/clients", http.StatusSeeOther)
}

func (c *Client) del(w http.ResponseWriter, r *http.Request) {
	c.mu.Lock()
	var id string
	if err := ps.Scanf(r.URL.Path, &id); err != nil {
		http.Redirect(w, r, "/api/clients", http.StatusSeeOther)
	}
	delete(c.clients, id)
	c.mu.Unlock()
	http.Redirect(w, r, "/api/clients", http.StatusSeeOther)
}
