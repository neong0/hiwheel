package wheel

import (
	"log"
	"net/http"
)

type HandlerFunc func(*Context)

type RouterGroup struct {
	prefix     string
	parent     *RouterGroup
	middleware []HandlerFunc
	engine     *Engine
}

func (g *RouterGroup) Group(prefix string) *RouterGroup {
	eng := g.engine
	newGroup := &RouterGroup{
		prefix:     g.prefix + prefix,
		parent:     g.engine.RouterGroup,
		middleware: make([]HandlerFunc, 0),
		engine:     eng,
	}
	eng.groups = append(eng.groups, *newGroup)
	return newGroup
}

func (g *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := g.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	g.engine.route.addRoute(method, pattern, handler)
}

func (g *RouterGroup) GET(pattern string, handler HandlerFunc) {
	g.addRoute("GET", pattern, handler)
}

func (g *RouterGroup) POST(pattern string, handler HandlerFunc) {
	g.addRoute("POST", pattern, handler)
}

func (g *RouterGroup) DELETE(pattern string, fn HandlerFunc) {
	g.addRoute("DELETE", pattern, fn)
}

func (g *RouterGroup) PUT(pattern string, fn HandlerFunc) {
	g.addRoute("PUT", pattern, fn)
}

type Engine struct {
	route *router
	*RouterGroup
	groups []RouterGroup
}

func New() *Engine {
	engine := &Engine{route: newRouter()}
	engine.groups = make([]RouterGroup, 0)
	engine.RouterGroup = &RouterGroup{engine: engine}
	return engine
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
