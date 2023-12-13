package service

import (
	"task_management/model"
	"task_management/repository"
)

// provides business logic for the tasks
type TaskService struct {
	taskRepository *repository.TaskRepository
}

// create a new instance of TaskService 
func NewTaskService(repo *repository.TaskRepository) *TaskService {
	return &TaskService{taskRepository: repo}
}

// adds a new task
func (s *TaskService) AddTask(title string) {
	task := model.Task{Title: title}
	s.taskRepository.AddTask(task)
}

// returns all tasks
func (s *TaskService) GetTasks() []model.Task {
	return s.taskRepository.GetTasks()
}