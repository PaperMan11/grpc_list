package controller

import (
	"api-gateway/pb"
	"context"

	"github.com/gin-gonic/gin"
)

func TaskCreate(c *gin.Context) {
	var req pb.TaskRequest
	var ok bool
	var taskSrv pb.TaskServiceClient
	err := c.BindJSON(&req)
	if err != nil {
		ResponseWithJSON(c, 500, "failed", nil)
		return
	}
	req.UserId, ok = c.Keys["userId"].(int64)
	if !ok {
		ResponseWithJSON(c, 500, "failed", nil)
		return
	}
	taskSrv, ok = c.Keys["task"].(pb.TaskServiceClient)
	if !ok {
		ResponseWithJSON(c, 500, "failed", nil)
		return
	}
	resp, err := taskSrv.TaskCreate(context.TODO(), &req)
	if err != nil {
		ResponseWithJSON(c, 500, "failed", nil)
		return
	}
	ResponseWithJSON(c, 200, "create success", resp)
}

func TaskUpdate(c *gin.Context) {
	var req pb.TaskRequest
	err := c.BindJSON(&req)
	if err != nil {
		ResponseWithJSON(c, 500, "failed1", nil)
		return
	}
	taskSrv, ok := c.Keys["task"].(pb.TaskServiceClient)
	if !ok {
		ResponseWithJSON(c, 500, "failed2", nil)
		return
	}
	resp, err := taskSrv.TaskUpdate(context.TODO(), &req)
	if err != nil {
		ResponseWithJSON(c, 500, "failed3", nil)
		return
	}
	ResponseWithJSON(c, 200, "update success", resp)
}
