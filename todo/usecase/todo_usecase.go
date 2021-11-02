package usecase

import (
	"github.com/boomauakim/go-todo-clean-arch/domain"
)

type todoUseCase struct {
	todoRepo domain.TodoRepository
}

func NewTodoUseCase(repo domain.TodoRepository) domain.TodoUseCase {
	return &todoUseCase{todoRepo: repo}
}

func (t *todoUseCase) ListAllTodos() (todos []domain.Todo, err error) {
	todos, err = t.todoRepo.ListAllTodos()

	return todos, err
}

func (t *todoUseCase) RetrieveTodo(id string) (todo domain.Todo, err error) {
	todo, err = t.todoRepo.RetrieveTodo(id)

	return todo, err
}
