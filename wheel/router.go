package wheel

import (
	"fmt"
	"net/http"
	"strings"
)

type H map[string]string

type router struct {
	RouterTree map[string]*routeNode
	Handlers   map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		Handlers:   make(map[string]HandlerFunc),
		RouterTree: make(map[string]*routeNode),
	}
}

func (r *router) addRoute(method string, pattern string, fn HandlerFunc) {
	pathSli := parsePattern(pattern)
	key := method + "-" + pattern
	if _, ok := r.RouterTree[method]; !ok {
		r.RouterTree[method] = &routeNode{}
	}
	r.RouterTree[method].insert(pattern, pathSli, 0)
	r.Handlers[key] = fn
}

func (r *router) Handle(c *Context) {
	n, param := r.getRoute(c.Method, c.Path)
	if n != nil {
		c.Params = param
		key := c.Method + "-" + c.Path
		r.Handlers[key](c)
		fmt.Println(n)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}

}

func (r *router) getRoute(method string, path string) (*routeNode, map[string]string) {
	searchPath := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.RouterTree[method]

	if !ok {
		return nil, nil
	}

	rn := root.search(searchPath, 0)

	if rn != nil {
		parts := parsePattern(rn.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchPath[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchPath[index:], "/")
				break
			}
		}
		return rn, params
	}
	return nil, nil
}

func parsePattern(pattern string) []string {
	temp := strings.Split(pattern, "/")
	res := make([]string, 0)
	for _, t := range temp {
		if t != "" {
			res = append(res, t)
			if t[0] == '*' {
				return res
			}
		}
	}
	return res
}
