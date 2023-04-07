package main

import (
	"learn/mvc/config"
	"learn/mvc/route"
)

func main() {
	config.InitDB()

	e := route.New()

	e.Start(":10000")
}
