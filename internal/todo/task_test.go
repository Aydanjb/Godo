package todo

import (
	"testing"
	"time"
)

func TestMarkTodo(t *testing.T) {
	tl := &TaskList{}
	task := tl.CreateTask("Wash dishes")
	task.MarkInProgress()

	t.Run("MarkTodo successfully marks task", func(t *testing.T) {
		task.MarkTodo()
		if task.Status != Todo {
			t.Errorf("expected task status to be %d, got %d", Todo, task.Status)
		}

	})
}

func TestMarkInProgress(t *testing.T) {
	tl := &TaskList{}
	task := tl.CreateTask("Wash dishes")

	before := time.Now()
	task.MarkInProgress()
	after := time.Now()

	t.Run("Status is In Progress", func(t *testing.T) {
		if task.Status != InProgress {
			t.Errorf("expected status %v, got %v", InProgress, task.Status)
		}
	})

	t.Run("UpdatedAt is in valid range", func(t *testing.T) {
		if task.UpdatedAt.Before(before) || task.UpdatedAt.After(after) {
			t.Errorf("expected between %v and %v, got %v", before, after, task.UpdatedAt)
		}
	})
}

func TestMarkDone(t *testing.T) {
	tl := &TaskList{}
	task := tl.CreateTask("Wash dishes")

	before := time.Now()
	task.MarkDone()
	after := time.Now()

	t.Run("Status is Done", func(t *testing.T) {
		if task.Status != Done {
			t.Errorf("expected status %v, got %v", Done, task.Status)
		}
	})

	t.Run("UpdatedAt is in valid range", func(t *testing.T) {
		if task.UpdatedAt.Before(before) || task.UpdatedAt.After(after) {
			t.Errorf("expected between %v and %v, got %v", before, after, task.UpdatedAt)
		}
	})

}
