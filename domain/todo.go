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

type UpdateTodo struct {
	Title     string `validate:"omitempty"`
	Completed *bool
}

type TodoUseCase interface {
	ListAllTodos() (t []Todo, err error)
	RetrieveTodo(id string) (t Todo, err error)
	CreateTodo(td *CreateTodo) (err error)
	UpdateTodo(id string, tu *UpdateTodo) (err error)
	DeleteTodo(id string) (err error)
}

type TodoRepository interface {
	ListAllTodos() (t []Todo, err error)
	RetrieveTodo(id string) (t Todo, err error)
	CreateTodo(ct *CreateTodo) (err error)
	UpdateTodo(id string, td *Todo, ut *UpdateTodo) (err error)
	DeleteTodo(id string) (err error)
}
