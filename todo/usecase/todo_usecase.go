package usecase

import (
	"github.com/boomauakim/go-todo-clean-arch/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (t *todoUseCase) CreateTodo(td *domain.CreateTodo) (err error) {
	err = t.todoRepo.CreateTodo(td)

	return err
}

func (t *todoUseCase) UpdateTodo(id string, tu *domain.UpdateTodo) (err error) {
	td, err := t.RetrieveTodo(id)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return domain.ErrNotFound
		}
		return err
	}

	err = t.todoRepo.UpdateTodo(id, &td, tu)

	return err
}
