package controller

import "github.com/gin-gonic/gin"

func ResponseWithJSON(ctx *gin.Context, code int, msg string, data interface{}) {
	ctx.JSON(code, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

type Response struct {
	Status uint        `json:"Status"`
	Data   interface{} `json:"Data"`
	Msg    string      `json:"Msg"`
}

//TokenData 带有token的Data结构
type TokenData struct {
	User  interface{} `json:"User"`
	Token string      `json:"Token"`
}
