package main

import (
	"gdd"
	"net/http"
)

func main() {
	r := gdd.New()
	r.GET("/", func(c *gdd.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	r.GET("/hello", func(c *gdd.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *gdd.Context) {
		c.JSON(http.StatusOK, gdd.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}
