package service

import (
	"context"

	"github.com/chelium/golang-todo-example/todo/pkg/io"
)

// TodoService describes the service.
type TodoService interface {
	Get(ctx context.Context) (t []io.Todo, err error)
	Add(ctx context.Context, todo io.Todo) (t io.Todo, err error)
	SetComplete(ctx context.Context, id string) (err error)
	RemoveComplete(ctx context.Context, id string) (err error)
	Delete(ctx context.Context, id string) (err error)
}

type basicTodoService struct{}

func (b *basicTodoService) Get(ctx context.Context) (t []io.Todo, err error) {
	// TODO implement the business logic of Get
	return t, err
}
func (b *basicTodoService) Add(ctx context.Context, todo io.Todo) (t io.Todo, err error) {
	// TODO implement the business logic of Add
	return t, err
}
func (b *basicTodoService) SetComplete(ctx context.Context, id string) (err error) {
	// TODO implement the business logic of SetComplete
	return err
}
func (b *basicTodoService) RemoveComplete(ctx context.Context, id string) (err error) {
	// TODO implement the business logic of RemoveComplete
	return err
}
func (b *basicTodoService) Delete(ctx context.Context, id string) (err error) {
	// TODO implement the business logic of Delete
	return err
}

// NewBasicTodoService returns a naive, stateless implementation of TodoService.
func NewBasicTodoService() TodoService {
	return &basicTodoService{}
}

// New returns a TodoService with all of the expected middleware wired in.
func New(middleware []Middleware) TodoService {
	var svc TodoService = NewBasicTodoService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
