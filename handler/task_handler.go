package handler

import (
	"fmt"
	"task_management/service"
)

// handles tasks from the console
type TaskHandler struct {
	taskService *service.TaskService
}

// creates a new instance of TaskHandler
func NewTaskHandler(service *service.TaskService) *TaskHandler {
	return &TaskHandler{taskService: service}
}

// starts the console application
func (h *TaskHandler) Run() {
	fmt.Println("Task management console App")
	for {
		fmt.Println("\n1. Add Task")
		fmt.Println("2. List Tasks")
		fmt.Println("3. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			h.addTask()
		case 2:
			h.listTasks()
		case 3:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Try again.")
		}
	}
}

func (h *TaskHandler) addTask() {
	var title string
	fmt.Print("Enter task title: ")
	fmt.Scanln(&title)

	h.taskService.AddTask(title)
	fmt.Println("Task added successfully!")
}

func (h *TaskHandler) listTasks() {
	tasks := h.taskService.GetTasks()
	fmt.Println("Tasks:")
	for _, task := range tasks {
		fmt.Printf("%d. %s\n", task.ID, task.Title)
	}
}