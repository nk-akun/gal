package gal

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Context ...
type Context struct {
	Writer http.ResponseWriter
	Req    *http.Request

	// Params store the parameters in path such as "/:name"
	Params     map[string]string
	Method     string
	Path       string
	StatusCode int

	middlewareIndex int

	// middleware functions
	handlers []HandleFunc
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer:          w,
		Req:             req,
		Method:          req.Method,
		Path:            req.URL.Path,
		middlewareIndex: -1,
	}
}

//Status set response status
func (c *Context) Status(statusCode int) {
	c.StatusCode = statusCode
	c.Writer.WriteHeader(statusCode)
}

//SetHeader ...
func (c *Context) SetHeader(key, value string) {
	c.Writer.Header().Add(key, value)
}

// Param get the para from context like /get/:name
func (c *Context) Param(key string) string {
	return c.Params[key]
}

//PostForm ...
func (c *Context) PostForm(key string) {
	c.Req.FormValue(key)
}

// Query ...
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// Next is the func from one middleware function to another
func (c *Context) Next() {
	for {
		c.middlewareIndex++
		if c.middlewareIndex >= len(c.handlers) {
			break
		}
		c.handlers[c.middlewareIndex](c)
	}
}

// String ...
func (c *Context) String(statusCode int, format string, values ...interface{}) {
	c.Status(statusCode)
	c.SetHeader("Content-Type", "text/plain")
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// JSON ...
func (c *Context) JSON(statusCode int, obj interface{}) {
	c.Status(statusCode)
	c.SetHeader("Content-Type", "application/json")
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

// Data ...
func (c *Context) Data(statusCode int, data []byte) {
	c.Status(statusCode)
	c.Writer.Write(data)
}

// HTML ...
func (c *Context) HTML(statusCode int, html string) {
	c.Status(statusCode)
	c.SetHeader("Content-Type", "text/html")
	c.Writer.Write([]byte(html))
}
