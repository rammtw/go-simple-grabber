package main

import (
	"lesson/news"
	"lesson/router"
)

func main() {
	r := router.New()
	a := news.New()

	go a.Serve()
	r.Run()
}