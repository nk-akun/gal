package gal

import "strings"

// Router ...
type Router struct {
	roots    map[string]*node
	headlers map[string]HandleFunc
}

// newRouter generate new router
func newRouter() *Router {
	return &Router{
		roots:    make(map[string]*node),
		headlers: make(map[string]HandleFunc),
	}
}

// parsePattern split pattern using "/"
func parsePattern(pattern string) (parts []string) {
	originParts := strings.Split(pattern, "/")
	parts = make([]string, 0)
	for _, part := range originParts {
		if part != "" {
			parts = append(parts, part)
			if part[0] == '*' {
				break
			}
		}
	}
	return parts
}

// getRoute get the real pattern of path and the parameters if path
func (router *Router) getRoute(method string, path string) (n *node, params map[string]string) {
	root, ok := router.roots[method]
	if !ok {
		return nil, nil
	}
	searchParts := parsePattern(path)
	n = root.search(searchParts, 0)
	if n == nil {
		return nil, nil
	}

	params = make(map[string]string)
	parts := parsePattern(n.pattern)
	for i, part := range parts {
		if part[0] == ':' {
			params[part[1:]] = searchParts[i]
		}
		if part[0] == '*' && len(part) > 1 {
			params[part[1:]] = strings.Join(searchParts[i:], "/")
			break
		}
	}
	return n, params
}

// addRoute add func for url
func (router *Router) addRoute(method, pattern string, hanler HandleFunc) {
	key := method + "-" + pattern
	router.headlers[key] = hanler

	parts := parsePattern(pattern)
	root, ok := router.roots[method]
	if !ok {
		root = &node{}
		router.roots[method] = root
	}
	root.insert(pattern, parts, 0)
}

func (router *Router) handle(c *Context) {
	n, params := router.getRoute(c.Method, c.Path)

	if n != nil {
		c.Params = params
		key := c.Method + "-" + n.pattern
		router.headlers[key](c)
	} else {
		c.Data(404, []byte("404 not found!"))
	}
}
