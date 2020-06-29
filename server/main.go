package main

import (
	"fmt"
	"html/template"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

type Task struct {
	Name string
	Done bool
}

func (t Task) Create() {
	println(t.Name)
}

func main() {

	http.HandleFunc("/hello", HelloHandler)
	http.HandleFunc("/tasks", TasksHandler)
	http.ListenAndServe(":8887", nil)

}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	id := uuid.NewV4().String()
	fmt.Fprintf(w, "Hello world, %v", id)
}

func TasksHandler(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("server/tasks.html"))

	tasks := []Task{
		{"Tarefa 1", false},
		{"Tarefa 2", true},
	}

	x := Task{"Tarefa 3", false}

	x.Create()

	tmpl.Execute(w, tasks)
}
