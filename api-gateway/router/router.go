package router

import (
	"api-gateway/controller"
	"api-gateway/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter(service ...interface{}) *gin.Engine {
	e := gin.Default()
	e.Use(middleware.SetSevice(service...))
	v1 := e.Group("/api/v1")
	{
		v1.POST("/user/login", controller.UserLogin)
		v1.POST("/user/register", controller.UserRegister)
	}
	v2 := v1.Group("/")
	v2.Use(middleware.JWT())
	{
		v2.POST("task", controller.TaskCreate)
		v2.PUT("task", controller.TaskUpdate)
	}
	return e
}
