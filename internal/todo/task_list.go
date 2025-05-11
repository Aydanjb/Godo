package todo

import (
	"fmt"
	"time"
)

type TaskList struct {
	Tasks  []*Task
	NextID int
}

// CreateTask returns a new Task with the given description and the current timestamp.
// The task is marked as Todo by default.
func (tl *TaskList) CreateTask(description string) *Task {
	task := Task{
		ID:          tl.NextID,
		Description: description,
		Status:      Todo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	tl.Tasks = append(tl.Tasks, &task)
	tl.NextID++
	return &task
}

func (tl *TaskList) GetTaskByID(id int) (*Task, int, error) {
	for i, task := range tl.Tasks {
		if task.ID == id {
			return task, i, nil
		}
	}

	return nil, -1, fmt.Errorf("task %d not found", id)
}

func (tl *TaskList) UpdateTask(id int, description string) error {
	task, _, _ := tl.GetTaskByID(id)
	if task == nil {
		return fmt.Errorf("task with id %d not found", id)
	}

	task.Description = description
	task.UpdatedAt = time.Now()
	return nil
}

func (tl *TaskList) DeleteTask(id int) error {
	_, i, err := tl.GetTaskByID(id)
	if err != nil {
		return err
	}

	tl.Tasks = append(tl.Tasks[:i], tl.Tasks[i+1:]...)
	return nil
}
