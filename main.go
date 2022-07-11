package main

import (
	"net/http"

	"github.com/guaiu/myGee/web"
)

func main() {
	r := web.New()
	r.GET("/", func(c *web.Context) {
		c.HTML(http.StatusOK, "<h1>Hello myGee</h1>")
	})

	r.GET("/hello", func(c *web.Context) {
		// expect /hello?name=guaiu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *web.Context) {
		// expect /hello/guaiu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *web.Context) {
		c.JSON(http.StatusOK, web.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")
}
