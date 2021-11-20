package repository_test

import (
	"context"
	"testing"

	"github.com/boomauakim/go-todo-clean-arch/domain"
	"github.com/boomauakim/go-todo-clean-arch/infrastructure/firestore"
	"github.com/boomauakim/go-todo-clean-arch/todo/repository"
	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
	"google.golang.org/api/iterator"
)

func TestCreateTodo(t *testing.T) {
	client := firestore.NewFirestoreTestClient(context.Background())
	defer client.Close()

	todoRepo := repository.NewTodoRepository(client)
	err := todoRepo.CreateTodo(&domain.CreateTodo{
		Title: "Test",
	})

	assert.NoError(t, err)
}

func TestUpdateTodo(t *testing.T) {
	ctx := context.Background()
	client := firestore.NewFirestoreTestClient(ctx)
	defer client.Close()

	todoRepo := repository.NewTodoRepository(client)

	todos := make([]domain.Todo, 0)
	iter := client.Collection("todo").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		todo := domain.Todo{}
		mapstructure.Decode(doc.Data(), &todo)
		todo.ID = doc.Ref.ID
		todos = append(todos, todo)
	}

	completed := true
	err := todoRepo.UpdateTodo(todos[0].ID, &todos[0], &domain.UpdateTodo{
		Completed: &completed,
	})

	doc, err := client.Collection("todo").Doc(todos[0].ID).Get(ctx)
	todo := domain.Todo{}
	mapstructure.Decode(doc.Data(), &todo)
	todo.ID = doc.Ref.ID

	assert.NoError(t, err)
	assert.Equal(t, domain.Todo{
		ID:        todos[0].ID,
		Title:     todos[0].Title,
		Completed: true,
		CreatedAt: todos[0].CreatedAt,
	}, todo)
}

func TestListAllTodos(t *testing.T) {
	ctx := context.Background()
	client := firestore.NewFirestoreTestClient(ctx)
	defer client.Close()

	todoRepo := repository.NewTodoRepository(client)
	todos, err := todoRepo.ListAllTodos()

	assert.NoError(t, err)
	assert.NotEmpty(t, todos)
}

func TestRetrieveTodo(t *testing.T) {
	ctx := context.Background()
	client := firestore.NewFirestoreTestClient(ctx)
	defer client.Close()

	todoRepo := repository.NewTodoRepository(client)

	todos := make([]domain.Todo, 0)
	iter := client.Collection("todo").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		todo := domain.Todo{}
		mapstructure.Decode(doc.Data(), &todo)
		todo.ID = doc.Ref.ID
		todos = append(todos, todo)
	}

	todo, err := todoRepo.RetrieveTodo(todos[0].ID)

	assert.NoError(t, err)
	assert.Equal(t, todos[0], todo)
}

func TestDeleteTodo(t *testing.T) {
	ctx := context.Background()
	client := firestore.NewFirestoreTestClient(ctx)
	defer client.Close()

	todoRepo := repository.NewTodoRepository(client)

	todos := make([]domain.Todo, 0)
	iter := client.Collection("todo").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		todo := domain.Todo{}
		mapstructure.Decode(doc.Data(), &todo)
		todo.ID = doc.Ref.ID
		todos = append(todos, todo)
	}

	errDelete := todoRepo.DeleteTodo(todos[0].ID)

	_, err := client.Collection("todo").Doc(todos[0].ID).Get(ctx)

	assert.NoError(t, errDelete)
	assert.Error(t, err)
}
