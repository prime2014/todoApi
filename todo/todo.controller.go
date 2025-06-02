package todo

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// The controller should be bound to the entity
type TodoController struct {
	Db *gorm.DB
}

func (c *TodoController) CreateTodo(ctx *fiber.Ctx) error {
	todo := Todo{}

	if err := ctx.BodyParser(&todo); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input data!",
		})

	}

	if todo.Title == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Title is required",
		})
	}

	todo, err := todo.Create(c.Db)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(todo)

}

func (c *TodoController) GetAllToDo(ctx *fiber.Ctx) error {

	todo := Todo{}

	fmt.Println("Fetchin todos...")
	todoList, err := todo.FetchAll(c.Db)

	if err != nil {
		return fmt.Errorf("internal server error")
	}

	return ctx.Status(fiber.StatusOK).JSON(todoList)

}
