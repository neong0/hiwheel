package wheel

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"strings"
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

func (g *RouterGroup) Use(middlewares ...HandlerFunc) {
	g.middleware = append(g.middleware, middlewares...)
}

func (g *RouterGroup) createStaticHandler(relatePath string, fd http.FileSystem) HandlerFunc {
	absolutePath := path.Join(g.prefix, relatePath)
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fd))
	return func(c *Context) {
		file := c.Param("filepath")
		if _, err := fd.Open(file); err != nil {
			c.HTTPStatus = http.StatusNotFound
			return
		}
		fileServer.ServeHTTP(c.Writer, c.Req)
	}
}

func (g *RouterGroup) Static(relatePath string, root string) {
	handler := g.createStaticHandler(relatePath, http.Dir(root))
	urlPatt := path.Join(relatePath, "/*filepath")
	g.GET(urlPatt, handler)
}

type Engine struct {
	*RouterGroup
	route         *router
	groups        []RouterGroup
	htmlTemplates *template.Template
	funcMap       template.FuncMap
}

func New() *Engine {
	engine := &Engine{route: newRouter()}
	engine.groups = make([]RouterGroup, 0)
	engine.RouterGroup = &RouterGroup{engine: engine}
	return engine
}

func Default() *Engine {
	e := New()
	e.Use(Wheelogger(), Recovery())
	return e
}

func (e *Engine) Run(port string) {
	http.ListenAndServe(port, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var middlewares []HandlerFunc
	for _, group := range e.groups {
		if strings.HasPrefix(r.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middleware...)
		}
	}
	c := newContext(w, r)
	c.engine = e
	c.handlers = middlewares
	e.route.Handle(c)
}

// func (e *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
// 	e.route.addRoute(method, pattern, handler)
// }

func (e *Engine) SetFuncMap(funcMap template.FuncMap) {
	e.funcMap = funcMap
}

func (e *Engine) LoadHTMLGlob(pattern string) {
	e.htmlTemplates = template.Must(template.New("").Funcs(e.funcMap).ParseGlob(pattern))
}
