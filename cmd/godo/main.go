package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"Godo/internal/todo"
	"github.com/urfave/cli/v2"
)

func main() {
	ts := todo.JSONTaskStore{Filepath: "tasks.json"}
	tl, err := ts.Load()
	if err != nil {
		_ = fmt.Errorf("error loading tasks.json: %v", err)
	}

	app := &cli.App{
		Name:  "godo",
		Usage: "a simple task manager",
		Commands: []*cli.Command{
			{
				Name:  "add",
				Usage: "add a new task",
				Action: func(c *cli.Context) error {
					description := c.Args().First()
					if description == "" {
						return fmt.Errorf("please provide a task description")
					}

					tl.CreateTask(description)
					fmt.Println("Added task:", description)
					return nil
				},
			},
			{
				Name:  "update",
				Usage: "update a task",
				Action: func(c *cli.Context) error {
					id := c.Args().First()
					if id == "" {
						return fmt.Errorf("please provide a task id")
					}
					description := c.Args().Get(1)
					if description == "" {
						return fmt.Errorf("please provide a task description")
					}

					idInt, _ := strconv.Atoi(id)
					tl.UpdateTask(idInt, description)
					fmt.Println("Updated task:", id)
					return nil
				},
			},
			{
				Name:  "delete",
				Usage: "delete a task",
				Action: func(c *cli.Context) error {
					id := c.Args().First()
					if id == "" {
						return fmt.Errorf("please provide a task id")
					}

					idInt, _ := strconv.Atoi(id)
					tl.DeleteTask(idInt)
					fmt.Println("Deleted task:", id)
					return nil
				},
			},
			{
				Name:  "list",
				Usage: "list all tasks",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "status",
						Value: "filter by status: todo, done, in-progress",
					},
				},
				Action: func(c *cli.Context) error {
					status, err := todo.ParseStatus(c.String("status"))
					if err != nil {
						fmt.Errorf("error parsing task status: %v", err)
					}

					if len(tl.Tasks) == 0 {
						fmt.Println("No tasks found.")
						return nil
					}

					for _, task := range tl.Tasks {
						if status == -1 || task.Status == status {
							fmt.Printf("#%d [%s] %s\n", task.ID, task.Status.String(), task.Description)
						}
					}
					return nil
				},
			},
			{
				Name:  "mark-in-progress",
				Usage: "marks a task as in-progress",
				Action: func(c *cli.Context) error {
					id := c.Args().First()
					if id == "" {
						return fmt.Errorf("please provide a task id")
					}

					idInt, _ := strconv.Atoi(id)
					task, _, err := tl.GetTaskByID(idInt)
					if err != nil {
						return err
					}

					task.MarkInProgress()
					return nil
				},
			},
			{
				Name:  "mark-done",
				Usage: "marks a task as done",
				Action: func(c *cli.Context) error {
					id := c.Args().First()
					if id == "" {
						return fmt.Errorf("please provide a task id")
					}

					idInt, _ := strconv.Atoi(id)
					task, _, err := tl.GetTaskByID(idInt)
					if err != nil {
						return err
					}

					task.MarkDone()
					return nil
				},
			},
			{
				Name:  "mark-todo",
				Usage: "marks a task as todo",
				Action: func(c *cli.Context) error {
					id := c.Args().First()
					if id == "" {
						return fmt.Errorf("please provide a task id")
					}

					idInt, _ := strconv.Atoi(id)
					task, _, err := tl.GetTaskByID(idInt)
					if err != nil {
						return err
					}

					task.MarkTodo()
					return nil
				},
			},
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	err = ts.Save(tl)
	if err != nil {
		log.Fatal(err)
	}
}
