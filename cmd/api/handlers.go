package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/aydanjb/Godo/internal/todo"
)

type Repository struct {
	TaskStore todo.JSONTaskStore
	TaskList  *todo.TaskList
}

func (repo *Repository) getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(repo.TaskList.Tasks)
	if err != nil {
		http.Error(w, "Failed to encode tasks", http.StatusInternalServerError)
		return
	}
}

func (repo *Repository) getTasksByStatus(w http.ResponseWriter, r *http.Request) {
	status, err := todo.ParseStatus(r.PathValue("status"))
	if err != nil {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var found []todo.Task
	for _, task := range repo.TaskList.Tasks {
		if status == -1 || task.Status == status {
			found = append(found, *task)
		}
	}

	err = json.NewEncoder(w).Encode(found)
}

func (repo *Repository) postTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	type TaskInput struct {
		Description string `json:"description"`
	}

	var input TaskInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	task := repo.TaskList.CreateTask(input.Description)
	err = json.NewEncoder(w).Encode(task)
}

func (repo *Repository) patchTask(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	task, _, _ := repo.TaskList.GetTaskByID(id)
	if task == nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	type TaskUpdate struct {
		Description *string `json:"description,omitempty"`
		Status      *string `json:"status,omitempty"`
	}

	var update TaskUpdate
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if update.Description != nil {
		task.Description = *update.Description
	}

	if update.Status != nil {
		status, err := todo.ParseStatus(*update.Status)
		if err != nil {
			http.Error(w, "Invalid status value", http.StatusBadRequest)
			return
		}
		task.Status = status
	}

	task.UpdatedAt = time.Now()

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(task)
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}

func (repo *Repository) deleteTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 0 {
		http.NotFound(w, r)
		return
	}

	err = repo.TaskList.DeleteTask(id)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
