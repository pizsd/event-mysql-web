package units

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ApiResponse(c *gin.Context, code int, msg string, data interface{}) {

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}
