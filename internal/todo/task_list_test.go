package todo

import (
	"testing"
	"time"
)

func TestCreateTask(t *testing.T) {
	tl := &TaskList{}
	expectedDescription := "Wash dishes"

	before := time.Now()
	task := tl.CreateTask(expectedDescription)
	after := time.Now()

	t.Run("ID is correct", func(t *testing.T) {
		if task.ID != 0 {
			t.Errorf("expected id 0, got %d", task.ID)
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

func TestGetTaskByID(t *testing.T) {
	tl := &TaskList{}
	task := tl.CreateTask("Wash dishes")

	t.Run("GetTaskByID successfully returns task", func(t *testing.T) {
		actual, id, err := tl.GetTaskByID(0)
		if err != nil {
			t.Errorf("unexpected error getting task by ID: %s", err)
		}

		if id != task.ID {
			t.Errorf("expected id %d, got %d", task.ID, id)
		}

		if actual.Description != task.Description {
			t.Errorf("expected description '%s', got '%s'", task.Description, actual.Description)
		}
	})

	t.Run("GetTaskByID returns error if task does not exist", func(t *testing.T) {
		_, _, err := tl.GetTaskByID(1)
		if err == nil {
			t.Errorf("expected error getting task by ID, but got nil")
		}
	})
}

func TestDeleteTask(t *testing.T) {
	tl := &TaskList{}
	task := tl.CreateTask("Wash dishes")

	t.Run("DeleteTask successfully deletes task", func(t *testing.T) {
		err := tl.DeleteTask(task.ID)
		if err != nil {
			t.Errorf("unexpected error deleting task: %s", err)
		}

		if len(tl.Tasks) != 0 {
			t.Errorf("expected 0 tasks, got %d", len(tl.Tasks))
		}
	})

	t.Run("DeleteTask returns error if task does not exist", func(t *testing.T) {
		err := tl.DeleteTask(1)
		if err == nil {
			t.Errorf("expected error deleting task, but got nil")
		}
	})
}

func TestUpdateTask(t *testing.T) {
	tl := &TaskList{}

	task := tl.CreateTask("Wash dishes")

	t.Run("UpdateTask successfully updates task", func(t *testing.T) {
		expectedDescription := "Take out trash"

		before := time.Now()
		err := tl.UpdateTask(task.ID, expectedDescription)
		after := time.Now()

		if err != nil {
			t.Errorf("unexpected error updating task: %s", err)
		}

		if task.Description != expectedDescription {
			t.Errorf("expected description '%s', got '%s'", "Take out trash", task.Description)
		}

		if task.UpdatedAt.Before(before) || task.UpdatedAt.After(after) {
			t.Errorf("expected between %v and %v, got %v", before, after, task.UpdatedAt)
		}
	})

	t.Run("UpdateTask returns error if task does not exist", func(t *testing.T) {
		err := tl.UpdateTask(1, "Take out trash")
		if err == nil {
			t.Errorf("expected error updating task, but got nil")
		}
	})
}
