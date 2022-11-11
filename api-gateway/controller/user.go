package controller

import (
	"api-gateway/pb"
	"api-gateway/pkg/jwt"
	"context"

	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {
	var req pb.LoginRequest
	if err := c.Bind(&req); err != nil {
		ResponseWithJSON(c, 200, "invalid req param", nil)
		return
	}
	userSrv, ok := c.Keys["user"].(pb.UserServiceClient)
	if !ok {
		ResponseWithJSON(c, 500, "failed", nil)
		return
	}
	userResp, err := userSrv.UserLogin(context.TODO(), &req)
	if err != nil {
		ResponseWithJSON(c, 500, "failed", nil)
		return
	}
	token, err := jwt.GenerateToken(req.UserId)
	if err != nil {
		ResponseWithJSON(c, 500, "failed", nil)
		return
	}
	resp := Response{
		Status: 200,
		Msg:    "login success",
		Data: TokenData{
			User:  userResp,
			Token: token,
		},
	}
	c.JSON(200, resp)
}

func UserRegister(c *gin.Context) {
	var req pb.RegRequest
	if err := c.BindJSON(&req); err != nil {
		ResponseWithJSON(c, 200, "invalid req param", nil)
		return
	}
	userSrv, ok := c.Keys["user"].(pb.UserServiceClient)
	if !ok {
		ResponseWithJSON(c, 500, "failed", nil)
		return
	}
	userResp, err := userSrv.UserRegister(context.TODO(), &req)
	if err != nil {
		ResponseWithJSON(c, 500, "failed", nil)
		return
	}
	ResponseWithJSON(c, 200, "login success", userResp)
}
