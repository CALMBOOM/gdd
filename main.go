package main

import (
	"gdd"
	"net/http"
)

func main() {
	r := gdd.New()
	r.GET("/", func(c *gdd.Context) {
		c.HTML(http.StatusOK, "<h1>Hello gdd</h1>")
	})

	r.GET("/hello", func(c *gdd.Context) {
		// expect /hello?name=gddktutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name/", func(c *gdd.Context) {
		// expect /hello/gddktutu
		c.String(http.StatusOK, "hello %s,you're at %s\n", c.Param("name"), c.Path)
	})
	r.GET("/hello/:name/:age/", func(c *gdd.Context) {
		// expect /hello/gddktutu
		c.String(http.StatusOK, "hello %s,%s, you're at %s\n", c.Param("name"), c.Param("age"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *gdd.Context) {
		c.JSON(http.StatusOK, gdd.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")
}
