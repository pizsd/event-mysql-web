package middleware

import (
	units "gin-web/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

const AuthorizationHeader = "Bearer "

func authAdminHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if len(token) == 0 {
			units.ApiResponse(c, http.StatusUnauthorized, "Invalid token", nil)
		}
		token = token[len(AuthorizationHeader):]
		claims, err := units.ParseToken(token)
		if err != nil {
			units.ApiResponse(c, http.StatusUnauthorized, err.Error(), nil)
		}
		//isExst :=
		c.Next()
	}
}

func authApiHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if len(token) == 0 {
			units.ApiResponse(c, http.StatusUnauthorized, "Invalid token", nil)
		}
		token = token[len(AuthorizationHeader):]
		_, err := units.ParseToken(token)
		if err != nil {
			units.ApiResponse(c, http.StatusUnauthorized, err.Error(), nil)
		}
		c.Next()
	}
}
