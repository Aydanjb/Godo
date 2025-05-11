package todo

import "time"

type Status int

const (
	Todo Status = iota
	InProgress
	Done
)

type Task struct {
	ID          int
	Description string
	Status      Status
	CreatedAt   time.Time
	UpdatedAt   time.Time
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
