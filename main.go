package main

import (
    "fmt"
    "os"

    "github.com/urfave/cli"
)

var tasks []string

func main() {
    app := cli.NewApp()
    app.Name = "Task Manager CLI"
    app.Usage = "Manage your tasks via command line"

    app.Commands = []cli.Command{
        {
            Name:  "submit",
            Usage: "Submit a new task",
            Action: func(c *cli.Context) error {
                task := c.Args().First()
                tasks = append(tasks, task)
                fmt.Printf("Task \"%s\" submitted!\n", task)
                return nil
            },
        },
        {
            Name:  "status",
            Usage: "List all tasks",
            Action: func(c *cli.Context) error {
                fmt.Println("Tasks:")
                for i, task := range tasks {
                    fmt.Printf("%d. %s\n", i+1, task)
                }
                return nil
            },
        },
        {
            Name:  "delete",
            Usage: "Delete a task",
            Action: func(c *cli.Context) error {
                index := c.Args().First()
                if index == "" {
                    fmt.Println("Please specify the task index to delete.")
                } else {
                    taskIndex := 0
                    fmt.Sscan(index, &taskIndex)
                    if taskIndex > 0 && taskIndex <= len(tasks) {
                        deletedTask := tasks[taskIndex-1]
                        tasks = append(tasks[:taskIndex-1], tasks[taskIndex:]...)
                        fmt.Printf("Task \"%s\" deleted!\n", deletedTask)
                    } else {
                        fmt.Println("Invalid task index.")
                    }
                }
                return nil
            },
        },
    }

    app.Run(os.Args)
}
