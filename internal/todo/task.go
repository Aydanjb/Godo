package todo

import (
	"fmt"
	"strings"
	"time"
)

type Status int

const (
	Todo Status = iota
	InProgress
	Done
)

func (s Status) String() string {
	switch s {
	case Todo:
		return "Todo"
	case InProgress:
		return "InProgress"
	case Done:
		return "Done"
	default:
		return "Unknown"
	}
}

func ParseStatus(s string) (Status, error) {
	s = strings.ToLower(s)
	switch s {
	case "todo":
		return Todo, nil
	case "in-progress":
		return InProgress, nil
	case "done":
		return Done, nil
	default:
		return -1, fmt.Errorf("unknown status %s", s)
	}
}

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

func (t *Task) MarkTodo() {
	t.Status = Todo
	t.UpdatedAt = time.Now()
}

// MarkInProgress sets a task's status to InProgress and updates UpdatedAt to the current time.
func (t *Task) MarkInProgress() {
	t.Status = InProgress
	t.UpdatedAt = time.Now()
}

// MarkDone sets a task's status to Done and updates UpdatedAt to the current time.
func (t *Task) MarkDone() {
	t.Status = Done
	t.UpdatedAt = time.Now()
}
