package main

import (
	"github.com/gominima/cors"
	"github.com/gominima/minima"
)

func main() {
	app := minima.New()
	crs := cors.New()
	app.Get("/", func(res *minima.Response, req *minima.Request) {
		res.OK().Send("Hello World")
		res.CloseConn()
	})
	app.Use(crs.AllowAll())
	app.Listen(":3000")
}
