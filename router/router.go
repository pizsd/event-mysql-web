package router

import (
	"gin-web/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Route() *gin.Engine {
	r := gin.Default()
	admin := r.Group("/admin")
	{
		admin.POST("/login")
		admin.GET("/home")
	}

	api := r.Group("/api")
	{
		api.POST("/login")
		api.GET("/user")
		api.POST("/user", controller.User)
		api.GET("/home")
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "PONG"})
	})
	return r
}
