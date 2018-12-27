package service

import (
	"context"
	io "github.com/chelium/golang-todo-example/todo/pkg/io"
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(TodoService) TodoService

type loggingMiddleware struct {
	logger log.Logger
	next   TodoService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a TodoService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next TodoService) TodoService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Get(ctx context.Context) (t []io.Todo, err error) {
	defer func() {
		l.logger.Log("method", "Get", "t", t, "err", err)
	}()
	return l.next.Get(ctx)
}
func (l loggingMiddleware) Add(ctx context.Context, todo io.Todo) (t io.Todo, err error) {
	defer func() {
		l.logger.Log("method", "Add", "todo", todo, "t", t, "err", err)
	}()
	return l.next.Add(ctx, todo)
}
func (l loggingMiddleware) SetComplete(ctx context.Context, id string) (err error) {
	defer func() {
		l.logger.Log("method", "SetComplete", "id", id, "err", err)
	}()
	return l.next.SetComplete(ctx, id)
}
func (l loggingMiddleware) RemoveComplete(ctx context.Context, id string) (err error) {
	defer func() {
		l.logger.Log("method", "RemoveComplete", "id", id, "err", err)
	}()
	return l.next.RemoveComplete(ctx, id)
}
func (l loggingMiddleware) Delete(ctx context.Context, id string) (err error) {
	defer func() {
		l.logger.Log("method", "Delete", "id", id, "err", err)
	}()
	return l.next.Delete(ctx, id)
}
