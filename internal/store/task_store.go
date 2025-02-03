package store

import (
	"time"
)

type Task struct {
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Priority    int32     `json:"priority"`
	Id          int32     `json:"id"`
}

type TaskStore interface {
	CreateTask(title, description string, priority int32) (Task, error)
	GetTaskById(id int32) (Task, error)
	ListTasks() ([]Task, error)
	UpdateTask(id int32, title, description string, priority int32) (Task, error)
	DeleteTask(id int32) error
}
