package main

import (
	"fmt"
	"gdd"
	"net/http"
)

func main() {
	r := gdd.New()

	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})
	r.Run(":80")
}
