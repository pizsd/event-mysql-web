package main

import (
	"gin-web/router"
)

func main() {
	engine := router.Route()
	_ = engine.Run()

}
