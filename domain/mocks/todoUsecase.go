package mocks

import (
	"github.com/boomauakim/go-todo-clean-arch/domain"
	"github.com/stretchr/testify/mock"
)

type TodoUsecase struct {
	mock.Mock
}

func (m *TodoUsecase) ListAllTodos() (todos []domain.Todo, err error) {
	ret := m.Called()

	return ret.Get(0).([]domain.Todo), ret.Error(1)
}

func (m *TodoUsecase) RetrieveTodo(id string) (todo domain.Todo, err error) {
	ret := m.Called()

	return ret.Get(0).(domain.Todo), ret.Error(1)
}

func (m *TodoUsecase) CreateTodo(td *domain.CreateTodo) (err error) {
	ret := m.Called()

	return ret.Error(0)
}

func (m *TodoUsecase) UpdateTodo(id string, tu *domain.UpdateTodo) (err error) {
	ret := m.Called()

	return ret.Error(0)
}

func (m *TodoUsecase) DeleteTodo(id string) (err error) {
	ret := m.Called()

	return ret.Error(0)
}
