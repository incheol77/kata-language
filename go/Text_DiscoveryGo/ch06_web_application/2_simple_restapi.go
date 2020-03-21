package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type ID string

type Task struct {
	task  string
	title string
	done  bool
	due   *time.Time
}

type DataAccess interface {
	Get(id ID) (task.Task, error)
	Put(id ID, t task.Task) error
	Post(t task.Task) (ID, error)
	Delete(id ID) error
}

type MemoryDataAccess struct {
	tasks  map[ID]task.Task
	nextID int64
}

func NewMemoryDataAccess() DataAccess {
	return &MemoryDataAccess{
		tasks:  map[ID]task.Task{},
		nextID: int64(1),
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, world")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
