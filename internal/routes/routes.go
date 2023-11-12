package routes

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	model "github.com/Michael-Sjogren/gohtmx/internal/model"
	fiber "github.com/gofiber/fiber/v2"
)

func GetTodos(ctx *fiber.Ctx) error {
	todos, err := model.GetAllTodos(50)
	if err != nil {
		log.Println(err)
		ctx.Status(fiber.StatusBadRequest)
		return nil
	}
	tmpl := template.Must(template.ParseFiles("templates/todos.html"))

	err = tmpl.ExecuteTemplate(ctx.Response().BodyWriter(), "Todos", todos)

	if err != nil {
		log.Println(err)
		ctx.Status(fiber.StatusInternalServerError)
	}

	return nil

}

func HandleDeleteTodo(ctx *fiber.Ctx) error {
	id, err := strconv.ParseInt(ctx.Params("id", "-1"), 10, 64)

	if err != nil || id == -1 {
		log.Println("Failed parse item id", err)
		ctx.Status(fiber.StatusBadRequest)
		return nil
	}

	err = model.DeleteTodo(id)

	if err != nil {
		log.Println("Failed to delete item", err)
		ctx.Status(fiber.StatusBadRequest)
	}

	todos, err := model.GetAllTodos(-1)
	tmpl := template.Must(template.ParseFiles("templates/todos.html"))

	err = tmpl.Execute(ctx.Response().BodyWriter(), todos)

	if err != nil {
		log.Println("template execution failed", err)
		ctx.Status(fiber.StatusInternalServerError)

		return nil
	}

	return nil

}

func CreateTodo(ctx *fiber.Ctx) error {
	description := ctx.FormValue("description", "")

	todo, err := model.CreateTodo(description, false)

	log.Println("TODO:", todo)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return nil
	}

	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	err = tmpl.Execute(ctx.Response().BodyWriter(), nil)

	if err != nil {
		log.Println("template execution failed", err)
		ctx.Status(fiber.StatusInternalServerError)
		return nil
	}
	return nil

}

func UpdateTodo(ctx *fiber.Ctx) error {
	id, err := strconv.ParseInt(ctx.Params("id", "-1"), 10, 64)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return err
	}
	description := ctx.FormValue("description", "")
	if len(description) == 0 {
		ctx.Status(fiber.StatusBadRequest)
		return nil
	}

	err = model.UpdateTodo(id, description, false)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return err
	}
	err = ctx.Render("templates/todo-item.html", model.Todo{
		Id:          id,
		Description: description,
		IsDone:      0,
	})

	if err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return err
	}

	return nil
}

func index(ctx *fiber.Ctx) error {

	err := ctx.Render("templates/index.html", nil)

	if err != nil {
		log.Println(err)
		ctx.Status(http.StatusBadRequest)
		return nil
	}

	return nil

}

func SetupServerAndRun() {
	app := fiber.New(
		fiber.Config{
			// Views Layout is the global layout for all template render until override on Render function.
			ViewsLayout: "templates/index.html",
		},
	)

	app.Static("/", "./static/", fiber.Static{
		Compress: true,
	})
	app.Get("/", index)
	app.Put("/todos/:id", UpdateTodo)
	app.Delete("/todos/:id", HandleDeleteTodo)
	app.Post("/todos", CreateTodo)
	app.Get("/todos", GetTodos)
	app.Get("/todos/:id", GetTodos)

	app.Listen(":8081")
}
