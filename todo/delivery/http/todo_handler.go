package http

import (
	"github.com/boomauakim/go-todo-clean-arch/domain"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type TodoHandler struct {
	todoUC domain.TodoUseCase
}

type ListAllTodosResponse struct {
	Todos []domain.Todo `json:"todos"`
}

type TodoResponse struct {
	Todo domain.Todo `json:"todo"`
}

func NewTodoHandler(uc domain.TodoUseCase) *TodoHandler {
	return &TodoHandler{
		todoUC: uc,
	}
}

func (t *TodoHandler) ListAllTodos(c *fiber.Ctx) error {
	todos, err := t.todoUC.ListAllTodos()
	if err != nil {
		return err
	}

	resp := ListAllTodosResponse{
		Todos: todos,
	}

	return c.JSON(resp)
}

func (t *TodoHandler) RetrieveTodo(c *fiber.Ctx) error {
	id := c.Params("id")

	todo, err := t.todoUC.RetrieveTodo(id)
	if err != nil {
		return err
	}

	resp := TodoResponse{
		todo,
	}

	return c.JSON(resp)
}
