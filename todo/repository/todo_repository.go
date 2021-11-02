package repository

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/boomauakim/go-todo-clean-arch/domain"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

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

func (t *todoRepository) RetrieveTodo(id string) (todo domain.Todo, err error) {
	ctx := context.Background()

	doc, err := t.fs.Collection("todo").Doc(id).Get(ctx)
	if err != nil {
		zap.L().Error("failed to retrieve a todo", zap.Error(err))
		if status.Code(err) == codes.NotFound {
			return domain.Todo{}, domain.ErrNotFound
		}
		return domain.Todo{}, err
	}

	todo = domain.Todo{}
	mapstructure.Decode(doc.Data(), &todo)
	todo.ID = doc.Ref.ID

	return todo, nil
}

func (t *todoRepository) CreateTodo(td *domain.CreateTodo) (err error) {
	ctx := context.Background()

	_, _, err = t.fs.Collection("todo").Add(ctx, map[string]interface{}{
		"title":     td.Title,
		"completed": false,
		"createdAt": firestore.ServerTimestamp,
	})
	if err != nil {
		zap.L().Error("failed to create a todo", zap.Error(err))
		return err
	}

	return nil
}

func (t *todoRepository) UpdateTodo(id string, td *domain.Todo, tu *domain.UpdateTodo) (err error) {
	ctx := context.Background()

	title := td.Title
	if tu.Title != "" {
		title = tu.Title
	}
	completed := td.Completed
	if tu.Completed != nil && completed != *tu.Completed {
		completed = *tu.Completed
	}
	_, err = t.fs.Collection("todo").Doc(id).Update(ctx, []firestore.Update{
		{
			Path:  "title",
			Value: title,
		},
		{
			Path:  "completed",
			Value: completed,
		},
	})
	if err != nil {
		zap.L().Error("failed to update a todo", zap.Error(err))
		return err
	}

	return nil
}

func (t *todoRepository) DeleteTodo(id string) (err error) {
	ctx := context.Background()

	_, err = t.fs.Collection("todo").Doc(id).Delete(ctx)
	if err != nil {
		zap.L().Error("failed to delete a todo", zap.Error(err))
		return err
	}

	return nil
}
