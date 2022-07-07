package app

import (
	"github.com/aykutaslantas/mongodb-api/models"
	"github.com/aykutaslantas/mongodb-api/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"github.com/gofiber/fiber/v2"
)

type TodoHandler struct {
	Service services.TodoService
}

func (h TodoHandler) CreateTodo(c *fiber.Ctx) error {
	var todo models.Todo
	if err := c.BodyParser(&todo); //json'u struct'a çeviriyoruz
	 err != nil {
		return err
	}

	res, err := h.Service.TodoInsert(todo)
	if err != nil || res.Status == false {
		return err
	}

	return c.Status(http.StatusCreated).JSON(res)
}

func (h TodoHandler) GetAllTodo(c *fiber.Ctx) error {
	res, err := h.Service.TodoGetAll()
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(res)
}

func (h TodoHandler) DeleteTodo(c *fiber.Ctx) error {
	query := c.Params("id")
	cnv, _ := primitive.ObjectIDFromHex(query)

	result, err := h.Service.TodoDelete(cnv)

	if err != nil || result == false {
		return c.Status(http.StatusBadRequest).JSON(err)
	}
	return c.Status(http.StatusOK).JSON(result)
}