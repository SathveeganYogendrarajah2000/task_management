package repository

import "task_management/model"

// TaskRepository handles the persistence of tasks
type TaskRepository struct {
	tasks []model.Task
}

// creates new instance of TaskRepository
func NewTaskRepository() *TaskRepository {
	return &TaskRepository{tasks: make([]model.Task, 0)}
}

// adds a new task to the repository
func (repo *TaskRepository) AddTask(task model.Task) {
	task.ID = len(repo.tasks) + 1
	repo.tasks = append(repo.tasks, task)
}

// returns all task from the repository
func (repo *TaskRepository) GetTasks() []model.Task {
	return repo.tasks
}