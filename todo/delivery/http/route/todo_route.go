package route

import (
	delivery "github.com/boomauakim/go-todo-clean-arch/todo/delivery/http"
	"github.com/boomauakim/go-todo-clean-arch/todo/repository"
	"github.com/boomauakim/go-todo-clean-arch/todo/usecase"
	"github.com/gofiber/fiber/v2"
)

func SetupRouter(f *fiber.App) {
	todoRepo := repository.NewTodoRepository()
	todoUC := usecase.NewTodoUseCase(todoRepo)
	todoHandler := delivery.NewTodoHandler(todoUC)

	v1 := f.Group("/v1")
	v1.Get("/todos", todoHandler.ListAllTodos)
}
