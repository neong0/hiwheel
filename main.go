package main

import (
	"fmt"
	"hiwheel/wheel"
	"log"
	"net/http"
	"time"
)

func testMiddleware(c *wheel.Context) {
	t := time.Now()
	c.String(500, "Internal Server Error")
	log.Printf("[%d] %s in %v for group", c.HTTPStatus, c.Req.URL, time.Since(t))
}

func main() {
	fmt.Println("hello internet")
	e := wheel.New()
	e.Use(wheel.WheeLogger())
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
