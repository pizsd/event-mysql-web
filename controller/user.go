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

func Login(c *gin.Context) {
	var userForm models.User
	if err := c.ShouldBind(&userForm); err != nil {
		myutils.ApiResponse(c, 400, err.Error(), nil)
		return
	}
	user := models.GetUserByName(userForm.Name)
	if user.Password != userForm.Password {
		myutils.ApiResponse(c, 400, "incorrect username or password", nil)
	}
	token, err := myutils.GenerateToken(user.Name, user.Password, "api")
	if err != nil {
		myutils.ApiResponse(c, 400, "token generation failed", nil)
	}
	myutils.ApiResponse(c, 200, "success", token)
}
