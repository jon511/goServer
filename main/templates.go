package main

import(
	"net/http"
	"html/template"
)

type Todo struct {
	Title string
	Done bool
}

type TodoPageData struct {
	PageTitle string
	Todos []Todo
}

func handleTemplate(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../templates/layout.html"))

	data := TodoPageData{
		PageTitle: "my TODO list",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: false},
			{Title: "Task 3", Done: true},
		},
	}

	tmpl.Execute(w, data)
}