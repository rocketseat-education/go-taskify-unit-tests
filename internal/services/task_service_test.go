package services

import (
	"testing"
	"time"

	"github.com/lohanguedes/taskify/internal/store"
	"github.com/stretchr/testify/assert"
)

type MockTaskStore struct{}

func (mocktaskstore *MockTaskStore) CreateTask(title string, description string, priority int32) (store.Task, error) {
	return store.Task{
		Id:          1,
		Title:       title,
		Description: description,
		Priority:    priority,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

func (mocktaskstore *MockTaskStore) GetTaskById(id int32) (store.Task, error) {
	return store.Task{
		Id:          id,
		Title:       "Mock Test Task",
		Description: "Mock Test Description",
		Priority:    1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

func (mocktaskstore *MockTaskStore) ListTasks() ([]store.Task, error) {
	return []store.Task{
		{
			Id:          1,
			Title:       "Mock Test Task",
			Description: "Mock Test Description",
			Priority:    1,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Id:          2,
			Title:       "Mock Test Task2",
			Description: "Mock Test Description2",
			Priority:    1,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}, nil
}

func (mocktaskstore *MockTaskStore) UpdateTask(id int32, title string, description string, priority int32) (store.Task, error) {
	return store.Task{
		Id:          id,
		Title:       title,
		Description: description,
		Priority:    priority,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

func (mocktaskstore *MockTaskStore) DeleteTask(id int32) error {
	return nil
}

func TestCreateTask(t *testing.T) {
	// Arrange
	mockStore := MockTaskStore{}
	taskService := NewTaskService(&mockStore)

	// Act
	task, err := taskService.Store.CreateTask("Mock Test Task", "Mock Test Description", 1)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, "Mock Test Task", task.Title)
	assert.Equal(t, "Mock Test Description", task.Description)
	assert.Equal(t, int32(1), task.Priority)
}

func TestGetTask(t *testing.T) {
	mockStore := MockTaskStore{}
	taskService := NewTaskService(&mockStore)

	task, err := taskService.Store.GetTaskById(1)

	assert.NoError(t, err)
	assert.Equal(t, int32(1), task.Id)
	assert.Equal(t, "Mock Test Task", task.Title)
}

func TestListTasks(t *testing.T) {
	// Arrange
	mockStore := MockTaskStore{}
	taskService := NewTaskService(&mockStore)

	// Act
	tasks, err := taskService.Store.ListTasks()

	assert.NoError(t, err)
	assert.Len(t, tasks, 2)
	assert.Equal(t, "Mock Test Task", tasks[0].Title)
	assert.Equal(t, "Mock Test Task2", tasks[1].Title)
}
