package main

import (
	"task_management/handler"
	"task_management/repository"
	"task_management/service"
)

func main() {
	taskRepo := repository.NewTaskRepository()

	taskService := service.NewTaskService(taskRepo)

	taskHandler := handler.NewTaskHandler(taskService)

	taskHandler.Run()
}