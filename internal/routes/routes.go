package routes

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	model "github.com/Michael-Sjogren/gohtmx/internal/model"
	"github.com/gorilla/mux"
)

var router *mux.Router

func GetTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := model.GetAllTodos(50)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	tmpl := template.Must(template.ParseFiles("templates/todos.html"))

	err = tmpl.ExecuteTemplate(w, "Todos", todos)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)

	if err != nil {
		log.Println("Failed parse item id", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = model.DeleteTodo(id)

	if err != nil {
		log.Println("Failed to delete item", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	todos, err := model.GetAllTodos(-1)

	tmpl := template.Must(template.ParseFiles("templates/todos.html"))

	err = tmpl.Execute(w, todos)

	if err != nil {
		log.Println("template execution failed", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	todo, err := model.CreateTodo(r.Form.Get("description"), false)

	log.Println("TODO:", todo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	err = tmpl.Execute(w, nil)

	if err != nil {
		log.Println("template execution failed", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {

}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	err := tmpl.Execute(w, nil)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func SetupServerAndRun() {
	router = mux.NewRouter()
	router.HandleFunc("/", index)
	router.HandleFunc("/todos/{id}", UpdateTodo).Methods("PUT")
	router.HandleFunc("/todos/{id}", DeleteTodo).Methods("DELETE")
	router.HandleFunc("/todos", CreateTodo).Methods("POST")
	router.HandleFunc("/todos", GetTodos).Methods("GET")
	router.HandleFunc("/todos/{id}", GetTodos).Methods("GET")

	log.Fatal(http.ListenAndServe(":8081", router))
}
