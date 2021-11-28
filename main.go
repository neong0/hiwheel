package main

import (
	"fmt"
	"hiwheel/wheel"
	"log"
	"net/http"
)

func main() {
	fmt.Println("hello internet")
	log.SetFlags(log.Ltime | log.Lshortfile)
	e := wheel.New()
	e.Get("/", func(c *wheel.Context) {
		c.HTML(http.StatusOK, "<h1>hello world<h2>")
	})

	e.Get("/hello", func(c *wheel.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})
	e.Get("/hello/:name", func(c *wheel.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	e.Get("/assets/*filepath", func(c *wheel.Context) {
		c.Json(http.StatusOK, wheel.H{"filepath": c.Param("filepath")})
	})

	e.Run(":9999")
}
