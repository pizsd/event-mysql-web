package main

import (
	"gin-web/router"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	router.Route(engine)
	_ = engine.Run()
}
