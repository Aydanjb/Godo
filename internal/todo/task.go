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

// CreateTask returns a new Task with the given description and the current timestamp.
// The task is marked as Todo by default.
func CreateTask(description string) Task {
	return Task{
		ID:          1,
		Description: description,
		Status:      Todo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (t *Task) MarkInProgress() {
	t.Status = InProgress
	t.UpdatedAt = time.Now()
}

// MarkDone sets a task's status to Done and updates UpdatedAt to the current time.
func (t *Task) MarkDone() {
	t.Status = Done
	t.UpdatedAt = time.Now()
}
