package services

import (
	"github.com/lohanguedes/taskify/internal/store"
)

type TaskService struct {
	Store store.TaskStore
}

func NewTaskService(store store.TaskStore) *TaskService {
	return &TaskService{Store: store}
}

func (s *TaskService) CreateTask(title, description string, priority int32) (store.Task, error) {
	// Add business logic here
	task, err := s.Store.CreateTask(title, description, priority)
	if err != nil {
		return store.Task{}, err
	}
	return task, err
}

func (s *TaskService) GetTask(id int32) (store.Task, error) {
	task, err := s.Store.GetTaskById(id)
	if err != nil {
		return store.Task{}, err
	}

	return task, err
}

func (s *TaskService) UpdateTask(id int32, title, description string, priority int32) (store.Task, error) {
	task, err := s.Store.UpdateTask(id, title, description, priority)
	if err != nil {
		return store.Task{}, err
	}

	return task, err
}

func (s *TaskService) DeleteTask(id int32) error {
	// Add business logic here...
	return s.Store.DeleteTask(id)
}
