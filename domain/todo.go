package domain

type Todo struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type TodoUseCase interface {
	ListAllTodos() ([]Todo, error)
}

type TodoRepository interface {
	ListAllTodos() (todos []Todo, err error)
}
