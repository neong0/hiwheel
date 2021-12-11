package main

import (
	"fmt"
	"hiwheel/wheel"
	"net/http"
)

func main() {
	fmt.Println("hello internet")
	e := wheel.Default()
	// e.LoadHTMLGlob("templates/*")
	// e.Static("/assets", "./static")
	e.GET("/panic", func(c *wheel.Context) {
		names := []string{"geektutu"}
		c.String(http.StatusOK, names[100])
	})
	e.GET("/", func(c *wheel.Context) {
		c.String(http.StatusOK, "Hello web\n")
	})
	v1 := e.Group("/v1")
	{
		// v1.GET("/", func(c *wheel.Context) {
		// 	c.HTML(http.StatusOK, "css.tmpl", nil)
		// })

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
