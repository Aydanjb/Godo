package todo

import (
	"os"
	"strconv"
	"testing"
)

func TestSaveAndLoad(t *testing.T) {
	tl := TaskList{}

	f, err := os.CreateTemp("", "tasks_*.json")
	if err != nil {
		t.Fatal(err)
	}
	fp := f.Name()
	defer os.Remove(fp)
	ts := JSONTaskStore{Filepath: fp}

	t.Run("Save and Load works", func(t *testing.T) {
		for i := 0; i < 5; i++ {
			tl.CreateTask("Task #" + strconv.Itoa(i))
		}

		err := ts.Save(&tl)
		if err != nil {
			t.Fatalf("failed to save: %v", err)
		}

		loaded, err := ts.Load()
		if err != nil {
			t.Fatalf("failed to load: %v", err)
		}

		if len(loaded.Tasks) != 5 {
			t.Fatalf("expected 5 tasks, got %d", len(loaded.Tasks))
		}

		for i, task := range loaded.Tasks {
			if loaded.Tasks[i].Description != "Task #"+strconv.Itoa(i) {
				t.Errorf("expected Task #%d, got %s", i, loaded.Tasks[0].Description)
			}

			if loaded.Tasks[i].ID != i {
				t.Errorf("expected %d, got %d", task.ID, loaded.Tasks[0].ID)
			}

			if loaded.Tasks[i].Status != Todo {
				t.Errorf("expected %d, got %d", task.Status, loaded.Tasks[0].Status)
			}
		}

	})

	t.Run("Load returns an empty TaskList if file doesn't exist", func(t *testing.T) {
		os.Remove(fp)
		loaded, err := ts.Load()
		if err != nil {
			t.Fatalf("failed to load: %v", err)
		}

		if len(loaded.Tasks) != 0 {
			t.Fatalf("expected 0 tasks, got %d", len(loaded.Tasks))
		}

		if loaded.NextID != 0 {
			t.Fatalf("expected %d, got %d", 0, loaded.NextID)
		}

	})
}
