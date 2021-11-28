package wheel

import (
	"net/http"
)

type HandlerFunc func(*Context)

type Engine struct {
	route *router
}

func New() *Engine {
	return &Engine{newRouter()}
}

func (e *Engine) Run(port string) {
	http.ListenAndServe(port, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContext(w, r)
	e.route.Handle(c)
}

func (e *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	e.route.addRoute(method, pattern, handler)
}

func (e *Engine) Get(pattern string, fn HandlerFunc) {
	e.addRoute("GET", pattern, fn)
}

func (e *Engine) POST(pattern string, fn HandlerFunc) {
	e.addRoute("POST", pattern, fn)
}

func (e *Engine) DELETE(pattern string, fn HandlerFunc) {
	e.addRoute("DELETE", pattern, fn)
}

func (e *Engine) PUT(pattern string, fn HandlerFunc) {
	e.addRoute("PUT", pattern, fn)
}
