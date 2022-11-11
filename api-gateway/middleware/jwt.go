package middleware

import (
	"api-gateway/pkg/jwt"
	"strings"
	"task/pkg/e"
	"time"

	"github.com/gin-gonic/gin"
)

// JWT token验证中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		code = 200
		token := strings.Split(c.Request.Header.Get("Authorization"), " ")[1]
		if token == "" {
			code = 404
		} else {
			claims, err := jwt.ParseToken(token)
			if err != nil {
				code = e.ErrorAuthCheckTokenFail
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ErrorAuthCheckTokenTimeout
			}
			c.Set("userId", claims.UserId)
		}
		if code != e.SUCCESS {
			c.JSON(200, gin.H{
				"status": code,
				"msg":    e.GetMsg(uint(code)),
				"data":   data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
