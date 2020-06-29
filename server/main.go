package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Task struct {
	Name string
	Done bool
}

func main() {

	http.HandleFunc("/hello", HelloHandler)
	http.HandleFunc("/tasks", TasksHandler)
	http.ListenAndServe(":8887", nil)

}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func TasksHandler(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("server/tasks.html"))

	tasks := []Task{
		{"Tarefa 1", false},
		{"Tarefa 2", true},
	}

	tmpl.Execute(w, tasks)
}
