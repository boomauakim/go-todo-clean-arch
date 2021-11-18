package mocks

import (
	"github.com/boomauakim/go-todo-clean-arch/domain"
	"github.com/stretchr/testify/mock"
)

type TodoRepository struct {
	mock.Mock
}

func (m *TodoRepository) ListAllTodos() (todos []domain.Todo, err error) {
	ret := m.Called()

	return ret.Get(0).([]domain.Todo), ret.Error(1)
}

func (m *TodoRepository) RetrieveTodo(id string) (todo domain.Todo, err error) {
	ret := m.Called()

	return ret.Get(0).(domain.Todo), ret.Error(1)
}

func (m *TodoRepository) CreateTodo(td *domain.CreateTodo) (err error) {
	ret := m.Called()

	return ret.Error(0)
}

func (m *TodoRepository) UpdateTodo(id string, td *domain.Todo, tu *domain.UpdateTodo) (err error) {
	ret := m.Called()

	return ret.Error(0)
}

func (m *TodoRepository) DeleteTodo(id string) (err error) {
	ret := m.Called()

	return ret.Error(0)
}
