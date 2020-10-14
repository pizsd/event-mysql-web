package controller

import (
	"gin-web/models"
	myutils "gin-web/utils"
	"github.com/gin-gonic/gin"
)

func User(c *gin.Context) {
	var userRegister models.User
	if err := c.ShouldBind(&userRegister); err != nil {
		myutils.ApiResponse(c, 400, err.Error(), nil)
		return
	}
	id, err := userRegister.Create()
	if err != nil {
		myutils.ApiResponse(c, 500, err.Error(), nil)
	} else {
		myutils.ApiResponse(c, 500, err.Error(), id)
	}
}
