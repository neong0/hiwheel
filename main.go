package main

import (
	"fmt"
	"hiwheel/wheel"
	"log"
	"net/http"
)

func main() {
	fmt.Println("hello internet")
	e := wheel.New()
	e.Get("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})

	e.Get("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})
	log.Fatal(http.ListenAndServe(":9999", e))
}
