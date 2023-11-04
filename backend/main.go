package main

import (
	"github.com/gin-gonic/gin"
	"task-management-backend/handlers"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
	"os"
	"fmt"
	"github.com/your-username/task-management-cli/handlers"
)

func main() {

	db, err := gorm.Open("postgres", "host=localhost username=cabontech Name=taskmanager sslmode=disable password=12345cabon")
    if err != nil {
        panic("Failed to connect to the database")
    }
    defer db.Close()


    args := os.Args[1:]
    if len(args) < 1 {
        fmt.Println("Usage: task <command>")
        return
}

command := args[0]
    switch command {
    case "status":
        handlers.Status(args[1:]) // Call the "status" handler with arguments
    case "submit":
        handlers.Submit(args[1:]) // Call the "submit" handler with arguments
    case "cancel":
        handlers.Cancel(args[1:]) // Call the "cancel" handler with arguments
    default:
        fmt.Printf("Unknown command: %s\n", command)
    }



    //db migration
    db.AutoMigrate(&models.Task{})

    
	r := gin.Default


	r.Run(":8080")
}
