package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gorilla/mux"
)

type Todo struct {
	ID    int
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

type TodoError struct {
	Message string
}

func (t *TodoError) Error() string {
	return t.Message
}

var dir, _ = os.Getwd()

var homeTemplate = filepath.Join(dir, "files", "home.html")
var notFoundTemplate = filepath.Join(dir, "files", "404.html")
var taskDetailsTemplate = filepath.Join(dir, "files", "task.html")

var data = TodoPageData{
	PageTitle: "My TODO list",
	Todos: []Todo{
		{ID: 1, Title: "Task 1", Done: false},
		{ID: 2, Title: "Task 2", Done: true},
		{ID: 3, Title: "Task 3", Done: true},
	},
}

func getTaskByID(id int) (*Todo, error) {

	for _, todo := range data.Todos {
		if todo.ID == id {
			return &todo, nil
		}
	}
	return nil, &TodoError{Message: "Todo not found"}
}

func home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(homeTemplate))
	fmt.Println("Showing home")
	tmpl.Execute(w, data)
	// fmt.Fprintln(w, "Welcome")
}

func taskDetailsHandler(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)

	taskID, err := strconv.Atoi(args["taskId"])
	if err != nil {
		//return 404

		tmpl := template.Must(template.ParseFiles(notFoundTemplate))

		tmpl.Execute(w, nil)
		return
	}

	task, err := getTaskByID(taskID)
	if err != nil {
		//return 404
		tmpl := template.Must(template.ParseFiles(notFoundTemplate))
		tmpl.Execute(w, nil)

		return
	}
	tmpl := template.Must(template.ParseFiles(taskDetailsTemplate))
	tmpl.Execute(w, task)
}
func main() {

	fmt.Println(homeTemplate)
	fmt.Println(notFoundTemplate)
	fmt.Println(taskDetailsTemplate)

	router := mux.NewRouter()
	router.HandleFunc("/", home)
	router.HandleFunc("/tasks", home)
	router.HandleFunc("/tasks/{taskId}", taskDetailsHandler)

	fmt.Println("Starting http server")
	http.ListenAndServe(":8090", router)

}
