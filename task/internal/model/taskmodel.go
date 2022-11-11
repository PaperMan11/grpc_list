package model

type Task struct {
	TaskId    int64  `db:"task_id"`
	UserId    int64  `db:"user_id"`
	Status    uint32 `db:"status"`
	Title     string `db:"title"`
	Content   string `db:"content"`
	StartTime uint32 `db:"start_time"`
	EndTime   uint32 `db:"end_time"`
}
