package gal

import (
	"net/http"
)

// HandleFunc the func to headle
type HandleFunc func(c *Context)

// RouterGroup ...
type RouterGroup struct {
	prefix      string
	middlewares []HandleFunc
	parent      *RouterGroup
	server      *Server
}

// Server engine of gal
type Server struct {
	*RouterGroup
	router *router
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

	//*Server implements ServeHTTP method
	return http.ListenAndServe(addr, server)
}

// ServerHTTP is the entrance
func (server *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	//generate new context once the request entries
	c := newContext(w, req)
	server.router.handle(c)
}
