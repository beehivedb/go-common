package main

import "github.com/beehivedb/go-common/router"

func main() {
	router.StaticPath("/", "static/")
	router.RegistryAction("/path", func(ctx router.Context) {
		ctx.Put("abc")
	})

	router.Run(":8080")
}

