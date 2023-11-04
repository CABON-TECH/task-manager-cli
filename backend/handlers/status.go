package handlers

import (
    "fmt"
    "strconv"
    "github.com/CABON-TECH/task-management-cli/models"
)

func Status(args []string) {
    if len(args) != 1 {
        fmt.Println("Usage: task status <task_id>")
        return
    }

    taskID, err := strconv.ParseUint(args[0], 10, 64)
    if err != nil {
        fmt.Printf("Invalid task ID: %v\n", err)
        return
    }

    task, err := models.GetTaskByID(uint(taskID))
    if err != nil {
        fmt.Printf("Failed to retrieve task: %v\n", err)
        return
    }

    fmt.Printf("Task ID: %d\nTitle: %s\nStatus: %s\nDescription: %s\n", task.ID, task.Title, task.Status, task.Description)
}
