package todo

import (
	"testing"
	"time"
)

func TestCreateTask(t *testing.T) {
	expectedDescription := "Wash dishes"
	before := time.Now()
	task := CreateTask(expectedDescription)
	after := time.Now()

	t.Run("ID is correct", func(t *testing.T) {
		if task.ID != 1 {
			t.Errorf("expected id 1, got %d", task.ID)
		}
	})

	t.Run("Description is set", func(t *testing.T) {
		if task.Description != expectedDescription {
			t.Errorf("expected '%s', got %s", expectedDescription, task.Description)
		}
	})

	t.Run("Status is Todo", func(t *testing.T) {
		if task.Status != Todo {
			t.Errorf("expected status %v, got %v", Todo, task.Status)
		}
	})

	t.Run("CreatedAt is in valid range", func(t *testing.T) {
		if task.CreatedAt.Before(before) || task.CreatedAt.After(after) {
			t.Errorf("expected between %v and %v, got %v", before, after, task.CreatedAt)
		}
	})

	t.Run("UpdatedAt is in valid range", func(t *testing.T) {
		if task.UpdatedAt.Before(before) || task.UpdatedAt.After(after) {
			t.Errorf("expected between %v and %v, got %v", before, after, task.UpdatedAt)
		}
	})
}

func TestMarkDone(t *testing.T) {
	expectedDescription := "Wash dishes"
	task := CreateTask(expectedDescription)

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
