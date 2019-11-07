package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"time"

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

func countTime() gal.HandleFunc {
	return func(c *gal.Context) {
		fmt.Println("process start~")
		last := time.Now().Unix()
		c.Next()
		now := time.Now().Unix()
		fmt.Println(last - now)
		fmt.Println("process end~")
	}
}

func getHeader() gal.HandleFunc {
	return func(c *gal.Context) {
		for key, value := range c.Req.Header {
			fmt.Println(key, ":", value)
		}
		c.Next()
	}
}

func testV3() {
	r := gal.New()

	r.Use(countTime())

	r.GET("/", func(c *gal.Context) {
		obj := map[string]interface{}{
			"name": "marathon",
			"age":  21,
		}
		c.JSON(200, obj)
	})

	v3 := r.Group("/v3")
	{
		v3.Use(getHeader())
		v3.GET("/hello/:name", func(c *gal.Context) {
			c.String(200, "How are you,%s? v2", c.Param("name"))
		})
		v3.GET("/assets/*file", func(c *gal.Context) {
			c.JSON(200, map[string]string{"filepath": c.Param("file")})
		})
	}

	r.SetFuncMap(template.FuncMap{
		"formatAsDate": func() {},
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")

	err := r.Run("127.0.0.1:8080")
	fmt.Println(err)
}

func main() {
	// testV1()
	// testV2()
	// testV3()
}
