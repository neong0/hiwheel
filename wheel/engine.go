package wheel

import (
	"log"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

type Engine struct {
	route map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{route: make(map[string]HandlerFunc)}
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path
	if f, ok := e.route[key]; ok {
		f(w, r)
		return
	}
	log.Printf("unrecognized path: %s", r.URL.Path)
}

func (e *Engine) Get(pattern string, fn HandlerFunc) {
	e.addRoute("GET", pattern, fn)
}
func (e *Engine) addRoute(method string, pattern string, fn HandlerFunc) {
	key := method + "-" + pattern
	e.route[key] = fn
}
