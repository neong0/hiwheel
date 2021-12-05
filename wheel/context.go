package wheel

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Context struct {
	Req    *http.Request
	Writer http.ResponseWriter

	Params     map[string]string
	Path       string
	Method     string
	HTTPStatus int

	handlers []HandlerFunc
	index    int
}

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Req:    r,
		Writer: w,
		Path:   r.URL.Path,
		Method: r.Method,
		Params: make(map[string]string),
		index:  -1,
	}
}

func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) SetHeader(key string, val string) {
	c.Writer.Header().Set(key, val)
}

func (c *Context) String(code int, format string, param ...interface{}) {
	c.SetHeader("Content/Type", "text/plain")
	c.HTTPStatus = code
	c.Writer.Write([]byte(fmt.Sprintf(format, param...)))
}

func (c *Context) Json(code int, object interface{}) {
	c.SetHeader("Content/Type", "application/json")
	c.HTTPStatus = code
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(object); err != nil {
		panic(err)
	}
}

func (c *Context) Data(code int, data []byte) {
	c.HTTPStatus = code
	c.Writer.Write(data)
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content/Type", "text/html")
	c.HTTPStatus = code
	c.Writer.Write([]byte(html))
}

func (c *Context) Param(key string) string {
	res := c.Params[key]
	return res
}

func (c *Context) Next() {
	c.index++
	s := len(c.handlers)
	for ; c.index < s; c.index++ {
		c.handlers[c.index](c)
	}
}
