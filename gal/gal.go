package gal

import (
	"html/template"
	"net/http"
	"path"
	"strings"
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
	router        *router
	groups        []*RouterGroup
	htmlTemplates *template.Template
	funcMap       template.FuncMap
}

// New generate a new server
func New() *Server {
	server := &Server{router: newRouter()}
	server.RouterGroup = &RouterGroup{server: server}
	server.groups = []*RouterGroup{server.RouterGroup}
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

// Use is the function to add a new middleware for a group
func (group *RouterGroup) Use(handler HandleFunc) {
	group.middlewares = append(group.middlewares, handler)
}

// SetFuncMap ...
func (server *Server) SetFuncMap(funcMap template.FuncMap) {
	server.funcMap = funcMap
}

// LoadHTMLGlob ...
func (server *Server) LoadHTMLGlob(pattern string) {
	server.htmlTemplates = template.Must(template.New("").Funcs(server.funcMap).ParseGlob(pattern))
}

// create static handler
func (group *RouterGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandleFunc {
	absolutePath := path.Join(group.prefix, relativePath)
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))
	return func(c *Context) {
		file := c.Param("filepath")
		// Check if file exists and/or if we have permission to access it
		if _, err := fs.Open(file); err != nil {
			c.Status(http.StatusNotFound)
			return
		}

		fileServer.ServeHTTP(c.Writer, c.Req)
	}
}

// Static serve static files
func (group *RouterGroup) Static(relativePath string, root string) {
	handler := group.createStaticHandler(relativePath, http.Dir(root))
	urlPattern := path.Join(relativePath, "/*filepath")
	// Register GET handlers
	group.GET(urlPattern, handler)
}

// ServerHTTP is the entrance
func (server *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	//generate new context once the request entries

	middlewares := []HandleFunc{}

	// get all middleware functions which match the req's URL
	for _, group := range server.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}

	c := newContext(w, req)
	c.handlers = middlewares
	c.server = server
	server.router.handle(c)
}
