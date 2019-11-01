package main

import (
	"fmt"
	"net/http"

	"github.com/nk-akun/gal/gal"
)

func main() {
	r := gal.New()
	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "hello")
	})

	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for key, v := range req.Header {
			fmt.Fprintf(w, "%s:%s\n", key, v)
		}
	})

	err := r.Run("127.0.0.1:8080")
	fmt.Println(err)
}
