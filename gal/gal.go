package gal

import (
	"net/http"
)

// HandleFunc the func to headle
type HandleFunc func(c *Context)

// Server engine of gal
type Server struct {
	router *Router
}

// New generate a new server
func New() *Server {
	return &Server{router: newRouter()}
}

// GET ...
func (server *Server) GET(pattern string, handler HandleFunc) {
	server.router.addRoute("GET", pattern, handler)
}

// POST ...
func (server *Server) POST(pattern string, handler HandleFunc) {
	server.router.addRoute("POST", pattern, handler)
}

// Run ...
func (server *Server) Run(addr string) error {
	return http.ListenAndServe(addr, server)
}

// ServerHTTP ...
func (server *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	server.router.handle(c)
}
