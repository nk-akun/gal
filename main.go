package main

import (
	"encoding/json"
	"fmt"

	"github.com/nk-akun/gal/gal"
)

func main() {
	r := gal.New()
	r.GET("/", func(c *gal.Context) {
		obj := map[string]interface{}{
			"name": "marathon",
			"age":  21,
		}
		ans, _ := json.Marshal(obj)
		c.Data(200, ans)
	})

	r.GET("/hello", func(c *gal.Context) {
		// c.String(200, "hello,your method is %s,your path is %s", c.Method, c.Path)
		c.HTML(200, "<h1>HELLO<h1>")
	})

	err := r.Run("127.0.0.1:8080")
	fmt.Println(err)
}
