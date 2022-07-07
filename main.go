package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/aykutaslantas/mongodb-api/configs"
	"github.com/aykutaslantas/mongodb-api/repository"
	"github.com/aykutaslantas/mongodb-api/services"
	"github.com/aykutaslantas/mongodb-api/app"
)

func main() {
	appRoute := fiber.New()
	configs.ConnectDB()
	dbClient := configs.GetCollection(configs.DB, "todos")

	TodoRepositoryDB := repository.NewTodoRepositoryDB(dbClient)

	td := app.TodoHandler{Service: services.NewTodoService(TodoRepositoryDB)}

	appRoute.Post("/api/todo", td.CreateTodo)
	appRoute.Get("/api/todos", td.GetAllTodo)
	appRoute.Delete("/api/todo/:id", td.DeleteTodo)
	appRoute.Listen(":8080")
}