package service

import (
	"context"

	"github.com/chelium/golang-todo-example/todo/pkg/db"
	"github.com/chelium/golang-todo-example/todo/pkg/io"
	"gopkg.in/mgo.v2/bson"
)

// TodoService describes the service.
type TodoService interface {
	Get(ctx context.Context) (t []io.Todo, erroror error)
	Add(ctx context.Context, todo io.Todo) (t io.Todo, erroror error)
	SetComplete(ctx context.Context, id string) (erroror error)
	RemoveComplete(ctx context.Context, id string) (erroror error)
	Delete(ctx context.Context, id string) (erroror error)
}

type basicTodoService struct{}

func (b *basicTodoService) Get(ctx context.Context) (t []io.Todo, error error) {
	session, err := db.GetMongoSession()
	if err != nil {
		return t, err
	}
	defer session.Close()
	c := session.DB("todo_app").C("todos")
	error = c.Find(nil).All(&t)
	return t, error
}
func (b *basicTodoService) Add(ctx context.Context, todo io.Todo) (t io.Todo, error error) {
	todo.Id = bson.NewObjectId()
	session, err := db.GetMongoSession()
	if err != nil {
		return t, err
	}
	defer session.Close()
	c := session.DB("todo_app").C("todos")
	error = c.Insert(&todo)
	return t, error
}
func (b *basicTodoService) SetComplete(ctx context.Context, id string) (error error) {
	session, err := db.GetMongoSession()
	if err != nil {
		return err
	}
	defer session.Close()
	c := session.DB("todo_app").C("todos")
	return c.Update(bson.M{"_id:": bson.ObjectIdHex(id)}, bson.M{"$set": bson.M{"complete": true}})
}
func (b *basicTodoService) RemoveComplete(ctx context.Context, id string) (error error) {
	session, err := db.GetMongoSession()
	if err != nil {
		return err
	}
	defer session.Close()
	c := session.DB("todo_app").C("todos")
	return c.Update(bson.M{"_id:": bson.ObjectIdHex(id)}, bson.M{"$set": bson.M{"complete": false}})
}
func (b *basicTodoService) Delete(ctx context.Context, id string) (error error) {
	// TODO implement the business logic of Delete
	return error
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
