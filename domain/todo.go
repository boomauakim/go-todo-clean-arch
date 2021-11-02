package domain

import "time"

type Todo struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateTodo struct {
	Title string `validate:"required"`
}

type TodoUseCase interface {
	ListAllTodos() (t []Todo, err error)
	RetrieveTodo(id string) (t Todo, err error)
	CreateTodo(td *CreateTodo) (err error)
}

type TodoRepository interface {
	ListAllTodos() (t []Todo, err error)
	RetrieveTodo(id string) (t Todo, err error)
	CreateTodo(td *CreateTodo) (err error)
}
