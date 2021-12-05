package wheel

import (
	"log"
	"time"
)

func Logger() HandlerFunc {
	return func(c *Context) {
		t := time.Now()
		c.Next()
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Printf("[%d] %s in %v", c.HTTPStatus, c.Req.URL, time.Since(t))
	}
}
