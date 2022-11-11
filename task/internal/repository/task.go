package repository

import (
	"log"
	"task/internal/model"
)

func CreateTask(task model.Task) error {
	sqlStr := `
		INSERT INTO task(task_id,user_id,status,title,content,start_time,end_time)
		VALUES(?, ?, ?, ?, ?, ?, ?)
	`
	if _, err := Db.Exec(sqlStr, task.TaskId, task.UserId, task.Status, task.Title, task.Content, task.StartTime, task.EndTime); err != nil {
		log.Println("create task failed")
		return err
	}

	return nil
}

func UpdateTask(task model.Task) error {
	sqlStr := `
		UPDATE task SET status=?,title=?,content=?,start_time=?,end_time=? 
		WHERE task_id=?
	`
	if _, err := Db.Exec(sqlStr, task.Status, task.Title, task.Content, task.StartTime, task.EndTime, task.TaskId); err != nil {
		log.Println("update task failed")
		return err
	}
	return nil
}

func QueryTaskById(taskId int64) (task *model.Task, err error) {
	task = new(model.Task)
	sqlStr := `SELECT * FROM task WHERE task_id = ?`
	err = Db.Get(task, sqlStr, taskId)
	if err != nil {
		log.Println("query task by id failed")
		return nil, err
	}
	return task, nil
}
