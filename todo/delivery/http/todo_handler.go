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

func (t *TodoHandler) CreateTodo(c *fiber.Ctx) error {
	todo := new(domain.CreateTodo)
	if err := c.BodyParser(todo); err != nil {
		return err
	}

	validate := validator.New()
	if err := validate.Struct(todo); err != nil {
		return fiber.NewError(fiber.ErrBadRequest.Code, err.Error())
	}

	err := t.todoUC.CreateTodo(todo)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).Send(nil)
}
