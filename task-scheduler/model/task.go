package model

type TaskMeta struct {
	TaskId   string `json: taskId`
	TaskType string `json: taskType`
	MaxRetry int    `json: maxRetry`
}

type Task struct {
	Id   string   `json: id`
	Meta TaskMeta `json: meta`
}