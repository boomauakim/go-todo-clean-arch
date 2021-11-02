package repository

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/boomauakim/go-todo-clean-arch/domain"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"google.golang.org/api/iterator"

type todoRepository struct {
	fs *firestore.Client
}

func NewTodoRepository(fs *firestore.Client) domain.TodoRepository {
	return &todoRepository{fs: fs}
}

func (t *todoRepository) ListAllTodos() (todos []domain.Todo, err error) {
	ctx := context.Background()

	todos = make([]domain.Todo, 0)
	iter := t.fs.Collection("todo").OrderBy("createdAt", firestore.Desc).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			zap.L().Error("failed to iter", zap.Error(err))
			return nil, err
		}
		todo := domain.Todo{}
		mapstructure.Decode(doc.Data(), &todo)
		todo.ID = doc.Ref.ID
		todos = append(todos, todo)
	}

	return todos, nil
}

		{
			ID:          "01FJVFD259PJ9RFJRGP0SPA963",
			Title:       "Todo 1",
			Description: "This first todo.",
		},
	}

	return todos, nil
}
