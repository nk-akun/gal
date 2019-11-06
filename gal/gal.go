package gal

import (
	"net/http"
)

// HandleFunc the func to headle
type HandleFunc func(c *Context)

// RouterGroup is struct for group router
type RouterGroup struct {
	prefix      string
	middlewares []HandleFunc

	//Design for multi routing
	parent *RouterGroup

	// Which server it belongs to
	server *Server
}

// Server is engine of gal
type Server struct {
	*RouterGroup
	router *router
	groups []*RouterGroup
}

// New generate a new server
func New() *Server {
	server := &Server{router: newRouter()}
	server.RouterGroup = &RouterGroup{server: server}
	server.groups = []*RouterGroup{}
	return server
}

// Group ...
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	server := group.server
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		server: server,
	}
	server.groups = append(server.groups, newGroup)
	return newGroup
}

// addRoute is the method for group router
func (group *RouterGroup) addRoute(method, pattern string, handler HandleFunc) {
	realPattern := group.prefix + pattern
	group.server.router.addRoute(method, realPattern, handler)
}

// GET ...
func (group *RouterGroup) GET(pattern string, handler HandleFunc) {
	group.addRoute("GET", pattern, handler)
}

// POST ...
func (group *RouterGroup) POST(pattern string, handler HandleFunc) {
	group.addRoute("POST", pattern, handler)
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
