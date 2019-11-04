package gal

// Router ...
type Router struct {
	headlers map[string]HandleFunc
}

func newRouter() *Router {
	return &Router{headlers: make(map[string]HandleFunc)}
}

func (router *Router) addRoute(method, pattern string, hanler HandleFunc) {
	key := method + "-" + pattern
	router.headlers[key] = hanler
}

func (router *Router) handle(c *Context) {
	k := c.Method + "-" + c.Path
	if headler, ok := router.headlers[k]; ok {
		headler(c)
	} else {
		c.Data(404, []byte("404 not found!"))
	}
}
