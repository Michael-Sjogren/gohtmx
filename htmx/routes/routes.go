package routes

import (
	"html/template"
	"htmxmain/htmx/model"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func APIGetTodos(w http.ResponseWriter, r *http.Request) {
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

func APIDeleteTodo(w http.ResponseWriter, r *http.Request) {

}

func APICreateTodo(w http.ResponseWriter, r *http.Request) {

}

func APIUpdateTodo(w http.ResponseWriter, r *http.Request) {

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

	mux.HandleFunc("/", index)
	mux.HandleFunc("/todos/{id}", APIUpdateTodo).Methods("PUT")
	mux.HandleFunc("/todos/{id}", APIDeleteTodo).Methods("DELETE")
	mux.HandleFunc("/todos", APICreateTodo).Methods("POST")
	mux.HandleFunc("/todos", APIGetTodos).Methods("GET")
	mux.HandleFunc("/todos/{id}", APIGetTodos).Methods("GET")

	log.Fatal(http.ListenAndServe(":8081", mux))
}
