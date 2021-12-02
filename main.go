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

	v1 := e.Group("/v1")
	{
		v1.GET("/", func(c *wheel.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
		})

		v1.GET("/hello", func(c *wheel.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := e.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *wheel.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *wheel.Context) {
			c.Json(http.StatusOK, wheel.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}

	e.Run(":9999")
}
