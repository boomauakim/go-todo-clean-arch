package repository

import "github.com/boomauakim/go-todo-clean-arch/domain"

type todoRepository struct{}

func NewTodoRepository() domain.TodoRepository {
	return &todoRepository{}
}

func (t *todoRepository) ListAllTodos() (todos []domain.Todo, err error) {
	todos = []domain.Todo{
		{
			ID:          "01FJVFD259PJ9RFJRGP0SPA963",
			Title:       "Todo 1",
			Description: "This first todo.",
		},
	}

	return todos, nil
}
