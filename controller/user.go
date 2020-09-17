package controller

import (
	"fmt"
	"gin-web/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func User(c *gin.Context) {
	var userRegister models.User
	if err := c.ShouldBind(&userRegister); err != nil {
		fmt.Println(userRegister)
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	id, err := userRegister.Create()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  err.Error(),
			"data": nil,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "ok",
			"data": id,
		})
	}

}
