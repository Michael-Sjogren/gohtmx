package routes

import (
	"html/template"
	"log"
	"net/http"

	"github.com/Michael-Sjogren/gohtmx/internal/model"

	"github.com/Michael-Sjogren/gohtmx/internal/middleware"

	"github.com/gorilla/mux"
)

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

}

func CreateTodo(w http.ResponseWriter, r *http.Request) {

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
	mux := mux.NewRouter()
	mux.Use(
		middleware.LoggingMiddleware,
	)
	mux.HandleFunc("/", index)
	mux.HandleFunc("/todos/{id}", UpdateTodo).Methods("PUT")
	mux.HandleFunc("/todos/{id}", DeleteTodo).Methods("DELETE")
	mux.HandleFunc("/todos", CreateTodo).Methods("POST")
	mux.HandleFunc("/todos", GetTodos).Methods("GET")
	mux.HandleFunc("/todos/{id}", GetTodos).Methods("GET")

	log.Fatal(http.ListenAndServe(":8081", mux))
}
