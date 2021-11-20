package usecase_test

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/boomauakim/go-todo-clean-arch/domain"
	"github.com/boomauakim/go-todo-clean-arch/domain/mocks"
	"github.com/boomauakim/go-todo-clean-arch/todo/usecase"
)

func TestListAllTodos(t *testing.T) {
	mockTodoRepo := new(mocks.TodoRepository)
	mockTodo := domain.Todo{
		ID:        "9lxa92ijz2xpWqWvPdrQ",
		Title:     "Test",
		Completed: false,
		CreatedAt: time.Now(),
	}
	mockListAllTodos := make([]domain.Todo, 0)
	mockListAllTodos = append(mockListAllTodos, mockTodo)

	t.Run("success", func(t *testing.T) {
		mockTodoRepo.On("ListAllTodos").Return(mockListAllTodos, nil).Once()

		todoUC := usecase.NewTodoUseCase(mockTodoRepo)
		todos, err := todoUC.ListAllTodos()

		assert.NoError(t, err)
		assert.Equal(t, mockListAllTodos, todos)
		assert.Len(t, todos, len(mockListAllTodos))

		mockTodoRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockTodoRepo.On("ListAllTodos").Return(make([]domain.Todo, 0), errors.New("Unexpected Error")).Once()

		todoUC := usecase.NewTodoUseCase(mockTodoRepo)
		todos, err := todoUC.ListAllTodos()

		assert.Empty(t, todos)
		assert.Error(t, err)

		mockTodoRepo.AssertExpectations(t)
	})
}

func TestRetrieveTodo(t *testing.T) {
	mockTodoRepo := new(mocks.TodoRepository)
	mockTodo := domain.Todo{
		ID:        "CLW7GOCG1vB6ChMawhLW",
		Title:     "Test",
		Completed: false,
		CreatedAt: time.Now(),
	}

	t.Run("success", func(t *testing.T) {
		mockTodoRepo.On("RetrieveTodo", mock.AnythingOfType("string")).Return(mockTodo, nil).Once()

		todoUC := usecase.NewTodoUseCase(mockTodoRepo)
		todo, err := todoUC.RetrieveTodo("CLW7GOCG1vB6ChMawhLW")

		assert.NoError(t, err)
		assert.NotNil(t, todo)

		mockTodoRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockTodoRepo.On("RetrieveTodo", mock.AnythingOfType("string")).Return(domain.Todo{}, errors.New("Unexpected Error")).Once()

		todoUC := usecase.NewTodoUseCase(mockTodoRepo)
		todo, err := todoUC.RetrieveTodo("CLW7GOCG1vB6ChMawhLW")

		assert.Empty(t, todo)
		assert.Error(t, err)

		mockTodoRepo.AssertExpectations(t)
	})
}

func TestCreateTodo(t *testing.T) {
	mockTodoRepo := new(mocks.TodoRepository)
	mockTodo := domain.CreateTodo{
		Title: "Test",
	}

	t.Run("success", func(t *testing.T) {
		mockTodoRepo.On("CreateTodo", mock.Anything).Return(nil).Once()

		todoUC := usecase.NewTodoUseCase(mockTodoRepo)
		err := todoUC.CreateTodo(&mockTodo)

		assert.NoError(t, err)

		mockTodoRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockTodoRepo.On("CreateTodo", mock.Anything).Return(errors.New("Unexpected Error")).Once()

		todoUC := usecase.NewTodoUseCase(mockTodoRepo)
		err := todoUC.CreateTodo(&mockTodo)

		assert.Error(t, err)

		mockTodoRepo.AssertExpectations(t)
	})
}

func TestUpdateTodo(t *testing.T) {
	mockTodoRepo := new(mocks.TodoRepository)
	mockTodo := domain.Todo{
		ID:        "CLW7GOCG1vB6ChMawhLW",
		Title:     "Test",
		Completed: false,
		CreatedAt: time.Now(),
	}
	completed := true
	mockTodoUpdate := domain.UpdateTodo{
		Completed: &completed,
	}

	t.Run("success", func(t *testing.T) {
		mockTodoRepo.On("RetrieveTodo", mock.AnythingOfType("string")).Return(mockTodo, nil).Once()
		mockTodoRepo.On("UpdateTodo", mock.AnythingOfType("string"), mock.Anything, mock.Anything).Return(nil).Once()

		todoUC := usecase.NewTodoUseCase(mockTodoRepo)
		err := todoUC.UpdateTodo("CLW7GOCG1vB6ChMawhLW", &mockTodoUpdate)

		assert.NoError(t, err)

		mockTodoRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockTodoRepo.On("RetrieveTodo", mock.AnythingOfType("string")).Return(mockTodo, nil).Once()
		mockTodoRepo.On("UpdateTodo", mock.AnythingOfType("string"), mock.Anything, mock.Anything).Return(errors.New("Unexpected Error")).Once()

		todoUC := usecase.NewTodoUseCase(mockTodoRepo)
		err := todoUC.UpdateTodo("CLW7GOCG1vB6ChMawhLW", &mockTodoUpdate)

		assert.Error(t, err)

		mockTodoRepo.AssertExpectations(t)
	})
}

func TestDeleteTodo(t *testing.T) {
	mockTodoRepo := new(mocks.TodoRepository)
	mockTodo := domain.Todo{
		ID:        "CLW7GOCG1vB6ChMawhLW",
		Title:     "Test",
		Completed: false,
		CreatedAt: time.Now(),
	}

	t.Run("success", func(t *testing.T) {
		mockTodoRepo.On("RetrieveTodo", mock.AnythingOfType("string")).Return(mockTodo, nil).Once()
		mockTodoRepo.On("DeleteTodo", mock.AnythingOfType("string")).Return(nil).Once()

		todoUC := usecase.NewTodoUseCase(mockTodoRepo)
		err := todoUC.DeleteTodo("CLW7GOCG1vB6ChMawhLW")

		assert.NoError(t, err)

		mockTodoRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockTodoRepo.On("RetrieveTodo", mock.AnythingOfType("string")).Return(mockTodo, nil).Once()
		mockTodoRepo.On("DeleteTodo", mock.AnythingOfType("string")).Return(errors.New("Unexpected Error")).Once()

		todoUC := usecase.NewTodoUseCase(mockTodoRepo)
		err := todoUC.DeleteTodo("CLW7GOCG1vB6ChMawhLW")

		assert.Error(t, err)

		mockTodoRepo.AssertExpectations(t)
	})
}
