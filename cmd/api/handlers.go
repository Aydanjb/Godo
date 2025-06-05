package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aydanjb/Godo/internal/todo"
)

type Repository struct {
	TaskStore todo.JSONTaskStore
	TaskList  *todo.TaskList
}

type TaskInput struct {
	Description string `json:"description"`
}

func (repo *Repository) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(repo.TaskList.Tasks)
	if err != nil {
		http.Error(w, "Failed to encode tasks", http.StatusInternalServerError)
		return
	}
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display all tasks..."))
}

func getTasksByStatus(w http.ResponseWriter, r *http.Request) {
	status, err := todo.ParseStatus(r.PathValue("status"))
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a all tasks where status == %s", status)
}

func (repo *Repository) postTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	var input TaskInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	task := repo.TaskList.CreateTask(input.Description)
	err = json.NewEncoder(w).Encode(task)
}

func patchStatus(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Mark a task as done, in-progress, todo..."))
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete a task"))
}
