package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/lohanguedes/taskify/internal/services"
)

type Application struct {
	Router      *chi.Mux
	TaskService services.TaskService
}
