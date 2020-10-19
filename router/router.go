package router

import (
	"gin-web/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Route() *gin.Engine {
	r := gin.Default()
	r.POST("/admin/login")
	admin := r.Group("/admin")
	//admin.Use()
	{
		admin.GET("/home")
	}
	r.POST("/api/login")
	api := r.Group("/api")
	{
		api.GET("/user")
		api.POST("/user", controller.User)
		api.GET("/home")
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "PONG"})
	})
	return r
}
