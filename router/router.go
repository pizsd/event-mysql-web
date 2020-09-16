package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Route(engine *gin.Engine) {
	admin := engine.Group("/admin")
	{
		admin.POST("/login")
		admin.GET("/home")
	}

	api := engine.Group("/api")
	{
		api.POST("/login")
		api.GET("/user")
		api.GET("/home")
	}

	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "PONG"})
	})
}
