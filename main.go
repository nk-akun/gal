package main

import (
	"encoding/json"
	"fmt"

	"github.com/nk-akun/gal/gal"
)

func testV1() {
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

func testV2() {
	r := gal.New()
	r.GET("/", func(c *gal.Context) {
		obj := map[string]interface{}{
			"name": "marathon",
			"age":  21,
		}
		c.JSON(200, obj)
	})

	r.GET("/hello", func(c *gal.Context) {
		c.HTML(200, "<h1>HELLO<h1>")
	})

	r.GET("/hello/:name", func(c *gal.Context) {
		c.String(200, "How are you,%s?", c.Param("name"))
	})

	r.GET("assets/*file", func(c *gal.Context) {
		c.JSON(200, map[string]string{"filepath": c.Param("file")})
	})

	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *gal.Context) {
			c.String(200, "How are you,%s? v2", c.Param("name"))
		})
		v2.GET("/assets/*file", func(c *gal.Context) {
			c.JSON(200, map[string]string{"filepath": c.Param("file")})
		})
	}

	err := r.Run("127.0.0.1:8080")
	fmt.Println(err)
}

func main() {
	// testV1()
	testV2()
}
