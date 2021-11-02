package domain

import "time"

type Todo struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

type TodoUseCase interface {
	ListAllTodos() (t []Todo, err error)
}

type TodoRepository interface {
	ListAllTodos() (t []Todo, err error)
}
