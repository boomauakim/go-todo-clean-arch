package http_test

import (
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/boomauakim/go-todo-clean-arch/domain"
	"github.com/boomauakim/go-todo-clean-arch/domain/mocks"
	delivery "github.com/boomauakim/go-todo-clean-arch/todo/delivery/http"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestListAllTodos(t *testing.T) {
	mockTodoUC := new(mocks.TodoUsecase)
	mockTodo := domain.Todo{
		ID:        "9lxa92ijz2xpWqWvPdrQ",
		Title:     "Test",
		Completed: false,
		CreatedAt: time.Now(),
	}
	mockListAllTodos := make([]domain.Todo, 0)
	mockListAllTodos = append(mockListAllTodos, mockTodo)

	mockTodoUC.On("ListAllTodos").Return(mockListAllTodos, nil)

	todoHandler := delivery.NewTodoHandler(mockTodoUC)

	app := fiber.New()
	app.Get("/todos", todoHandler.ListAllTodos)

	res, err := app.Test(httptest.NewRequest(fiber.MethodGet, "/todos", nil))

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, res.StatusCode)
}

func TestRetrieveTodo(t *testing.T) {
	mockTodoUC := new(mocks.TodoUsecase)
	mockTodo := domain.Todo{
		ID:        "CLW7GOCG1vB6ChMawhLW",
		Title:     "Test",
		Completed: false,
		CreatedAt: time.Now(),
	}

	mockTodoUC.On("RetrieveTodo", mock.AnythingOfType("string")).Return(mockTodo, nil)

	todoHandler := delivery.NewTodoHandler(mockTodoUC)

	app := fiber.New()
	app.Get("/todos/:id", todoHandler.RetrieveTodo)

	res, err := app.Test(httptest.NewRequest(fiber.MethodGet, "/todos/CLW7GOCG1vB6ChMawhLW", nil))

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, res.StatusCode)
}

func TestCreateTodo(t *testing.T) {
	mockTodoUC := new(mocks.TodoUsecase)
	mockTodo, _ := json.Marshal(domain.CreateTodo{
		Title: "Test",
	})

	mockTodoUC.On("CreateTodo", mock.Anything).Return(nil)

	todoHandler := delivery.NewTodoHandler(mockTodoUC)
	app := fiber.New()
	app.Post("/todos", todoHandler.CreateTodo)

	req := httptest.NewRequest(fiber.MethodPost, "/todos", strings.NewReader(string(mockTodo)))
	req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	res, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, res.StatusCode)
}

func TestUpdateTodo(t *testing.T) {
	mockTodoUC := new(mocks.TodoUsecase)
	completed := true
	mockTodo, _ := json.Marshal(domain.UpdateTodo{
		Completed: &completed,
	})

	mockTodoUC.On("UpdateTodo", mock.AnythingOfType("string"), mock.Anything).Return(nil)

	todoHandler := delivery.NewTodoHandler(mockTodoUC)
	app := fiber.New()
	app.Patch("/todos", todoHandler.UpdateTodo)

	req := httptest.NewRequest(fiber.MethodPatch, "/todos", strings.NewReader(string(mockTodo)))
	req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	res, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusNoContent, res.StatusCode)
}

func TestDeleteTodo(t *testing.T) {
	mockTodoUC := new(mocks.TodoUsecase)

	mockTodoUC.On("DeleteTodo", mock.AnythingOfType("string")).Return(nil)

	todoHandler := delivery.NewTodoHandler(mockTodoUC)

	app := fiber.New()
	app.Delete("/todos/:id", todoHandler.DeleteTodo)

	res, err := app.Test(httptest.NewRequest(fiber.MethodDelete, "/todos/CLW7GOCG1vB6ChMawhLW", nil))

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusNoContent, res.StatusCode)
}
