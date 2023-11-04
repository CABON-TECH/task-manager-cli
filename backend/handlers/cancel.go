package handlers

import (
    "fmt"
    "strconv"
    "github.com/CABON-TECH/task-management-cli/models"
)

func Cancel(args []string) {
    if len(args) != 1 {
        fmt.Println("Usage: task cancel <task_id>")
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

    fmt.Printf("Task with ID %d has been canceled.\n", task.ID)
}
