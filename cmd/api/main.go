package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/aydanjb/Godo/internal/todo"
)

func main() {
	ts := todo.JSONTaskStore{Filepath: "tasks.json"}
	tl, err := ts.Load()
	if err != nil {
		_ = fmt.Errorf("error loading tasks.json: %v", err)
		return
	}

	repo := &Repository{
		TaskList:  tl,
		TaskStore: ts,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /todo/{$}", repo.getTasks)
	mux.HandleFunc("GET /todo/{status}", repo.getTasksByStatus)
	mux.HandleFunc("POST /todo/create", repo.postTask)
	mux.HandleFunc("PATCH /todo/{id}", repo.patchTask)
	mux.HandleFunc("DELETE /todo/{id}", repo.deleteTask)

	server := &http.Server{
		Addr:    ":4000",
		Handler: mux,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go func() {
		log.Println("Server started on :4000")
		err := server.ListenAndServe()
		if err != nil {
			return
		}
	}()

	<-stop
	log.Println("Shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = server.Shutdown(ctx)
	if err != nil {
		log.Fatalf("Server Shutdown Failure: %+v", err)
	}

	err = repo.TaskStore.Save(repo.TaskList)
	if err != nil {
		log.Printf("Failed to save tasks: %v", err)
	} else {
		log.Printf("Saving %d tasks to %s\n", len(repo.TaskList.Tasks), ts.Filepath)
	}

	log.Println("Server shutdown gracefully")
}
