package middleware

import (
	"github.com/gin-gonic/gin"
)

// 接受服务实例，并存到gin.Key中
func SetSevice(service ...interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.Keys = make(map[string]interface{})
		// c.Keys["user"] = service[0]
		// // c.Keys["task"] = service[1]
		// c.Next()
		c.Set("user", service[0])
		c.Set("task", service[1])
		c.Next()
	}
}
