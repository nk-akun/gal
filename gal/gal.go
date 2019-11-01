package gal

import (
	"fmt"
	"net/http"
)

// HandleFunc the func to headle
type HandleFunc func(http.ResponseWriter, *http.Request)

// Server engine of gal
type Server struct {
	router map[string]HandleFunc
}

// New generate a new server
func New() *Server {
	return &Server{router: make(map[string]HandleFunc)}
}

func (server *Server) addRoute(method, pattern string, hanler HandleFunc) {
	key := method + "-" + pattern
	server.router[key] = hanler
}

// GET ...
func (server *Server) GET(pattern string, handler HandleFunc) {
	server.addRoute("GET", pattern, handler)
}

// POST ...
func (server *Server) POST(pattern string, handler HandleFunc) {
	server.addRoute("POST", pattern, handler)
}

// Run ...
func (server *Server) Run(addr string) error {
	return http.ListenAndServe(addr, server)
}

// ServerHTTP ...
func (server *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	k := req.Method + "-" + req.URL.Path
	if headler, ok := server.router[k]; ok {
		headler(w, req)
	} else {
		fmt.Fprintf(w, "404 not found!")
	}
}
