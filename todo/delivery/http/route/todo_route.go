package route

import (
	"github.com/boomauakim/go-todo-clean-arch/infrastructure/firestore"
	delivery "github.com/boomauakim/go-todo-clean-arch/todo/delivery/http"
	"github.com/boomauakim/go-todo-clean-arch/todo/repository"
	"github.com/boomauakim/go-todo-clean-arch/todo/usecase"
	"github.com/gofiber/fiber/v2"
)

func RegisterRouter(f *fiber.App) {
	client := firestore.NewFirestoreLocalClient()
	todoRepo := repository.NewTodoRepository(client)
	todoUC := usecase.NewTodoUseCase(todoRepo)
	todoHandler := delivery.NewTodoHandler(todoUC)

	v1 := f.Group("/v1")
	v1.Get("/todos", todoHandler.ListAllTodos)
	v1.Get("/todos/:id", todoHandler.RetrieveTodo)
	v1.Post("/todos", todoHandler.CreateTodo)
	v1.Patch("/todos/:id", todoHandler.UpdateTodo)
	v1.Delete("/todos/:id", todoHandler.DeleteTodo)
}
