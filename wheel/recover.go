package wheel

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
)

func Recovery() HandlerFunc {
	return func(c *Context) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("trace back")
				msg := fmt.Sprintf("%s", err)
				log.Printf("%s\n\n", trace(msg))
				c.String(http.StatusInternalServerError, "internet server error")
			}
		}()

		c.Next()
	}
}

func trace(err string) string {
	var temp [32]uintptr
	n := runtime.Callers(3, temp[:])

	var str strings.Builder
	str.WriteString(err + "\nTraceBack:")
	for _, pc := range temp[:n] {
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)
		str.WriteString(fmt.Sprintf("\n\t%s:%d", file, line))
	}
	return str.String()
}
