package service

import (
	"log"
	"task/internal/model"
	"task/internal/pb"
	"task/internal/repository"
	"task/pkg/snowflake"
)

func TaskCreate(req *pb.TaskRequest) error {
	var task model.Task
	task.UserId = req.UserId
	task.Status = req.Status
	task.Title = req.Title
	task.Content = req.Content
	task.StartTime = req.StartTime
	task.EndTime = req.EndTime
	// 生成task id
	task.TaskId = snowflake.GetID()
	if err := repository.CreateTask(task); err != nil {
		log.Println("CreateTask failed")
		return err
	}
	return nil
}

func TaskUpdate(req *pb.TaskRequest) error {
	var task model.Task
	task.UserId = req.UserId
	task.Status = req.Status
	task.Title = req.Title
	task.Content = req.Content
	task.StartTime = req.StartTime
	task.EndTime = req.EndTime
	task.TaskId = req.TaskId
	if err := repository.UpdateTask(task); err != nil {
		log.Println("UpdateTask failed")
		return err
	}
	return nil
}

func TaskShow(taskId int64) (task *model.Task, err error) {
	return repository.QueryTaskById(taskId)
}

func TaskDelete(req *pb.TaskRequest) {

}
