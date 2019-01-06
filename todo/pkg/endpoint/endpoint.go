package endpoint

import (
	context "context"
	io "github.com/chelium/golang-todo-example/todo/pkg/io"
	service "github.com/chelium/golang-todo-example/todo/pkg/service"
	endpoint "github.com/go-kit/kit/endpoint"
)

// GetRequest collects the request parameters for the Get method.
type GetRequest struct{}

// GetResponse collects the response parameters for the Get method.
type GetResponse struct {
	T   []io.Todo `json:"t"`
	Err error     `json:"err"`
}

// MakeGetEndpoint returns an endpoint that invokes Get on the service.
func MakeGetEndpoint(s service.TodoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		t, err := s.Get(ctx)
		return GetResponse{
			Err: err,
			T:   t,
		}, nil
	}
}

// Failed implements Failer.
func (r GetResponse) Failed() error {
	return r.Err
}

// AddRequest collects the request parameters for the Add method.
type AddRequest struct {
	Todo io.Todo `json:"todo"`
}

// AddResponse collects the response parameters for the Add method.
type AddResponse struct {
	T   io.Todo `json:"t"`
	Err error   `json:"err"`
}

// MakeAddEndpoint returns an endpoint that invokes Add on the service.
func MakeAddEndpoint(s service.TodoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddRequest)
		t, err := s.Add(ctx, req.Todo)
		return AddResponse{
			Err: err,
			T:   t,
		}, nil
	}
}

// Failed implements Failer.
func (r AddResponse) Failed() error {
	return r.Err
}

// SetCompleteRequest collects the request parameters for the SetComplete method.
type SetCompleteRequest struct {
	Id string `json:"id"`
}

// SetCompleteResponse collects the response parameters for the SetComplete method.
type SetCompleteResponse struct {
	Err error `json:"err"`
}

// MakeSetCompleteEndpoint returns an endpoint that invokes SetComplete on the service.
func MakeSetCompleteEndpoint(s service.TodoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SetCompleteRequest)
		err := s.SetComplete(ctx, req.Id)
		return SetCompleteResponse{Err: err}, nil
	}
}

// Failed implements Failer.
func (r SetCompleteResponse) Failed() error {
	return r.Err
}

// RemoveCompleteRequest collects the request parameters for the RemoveComplete method.
type RemoveCompleteRequest struct {
	Id string `json:"id"`
}

// RemoveCompleteResponse collects the response parameters for the RemoveComplete method.
type RemoveCompleteResponse struct {
	Err error `json:"err"`
}

// MakeRemoveCompleteEndpoint returns an endpoint that invokes RemoveComplete on the service.
func MakeRemoveCompleteEndpoint(s service.TodoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RemoveCompleteRequest)
		err := s.RemoveComplete(ctx, req.Id)
		return RemoveCompleteResponse{Err: err}, nil
	}
}

// Failed implements Failer.
func (r RemoveCompleteResponse) Failed() error {
	return r.Err
}

// DeleteRequest collects the request parameters for the Delete method.
type DeleteRequest struct {
	Id string `json:"id"`
}

// DeleteResponse collects the response parameters for the Delete method.
type DeleteResponse struct {
	Err error `json:"err"`
}

// MakeDeleteEndpoint returns an endpoint that invokes Delete on the service.
func MakeDeleteEndpoint(s service.TodoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteRequest)
		err := s.Delete(ctx, req.Id)
		return DeleteResponse{Err: err}, nil
	}
}

// Failed implements Failer.
func (r DeleteResponse) Failed() error {
	return r.Err
}

// Failer is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Get implements Service. Primarily useful in a client.
func (e Endpoints) Get(ctx context.Context) (t []io.Todo, err error) {
	request := GetRequest{}
	response, err := e.GetEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetResponse).T, response.(GetResponse).Err
}

// Add implements Service. Primarily useful in a client.
func (e Endpoints) Add(ctx context.Context, todo io.Todo) (t io.Todo, err error) {
	request := AddRequest{Todo: todo}
	response, err := e.AddEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(AddResponse).T, response.(AddResponse).Err
}

// SetComplete implements Service. Primarily useful in a client.
func (e Endpoints) SetComplete(ctx context.Context, id string) (err error) {
	request := SetCompleteRequest{Id: id}
	response, err := e.SetCompleteEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(SetCompleteResponse).Err
}

// RemoveComplete implements Service. Primarily useful in a client.
func (e Endpoints) RemoveComplete(ctx context.Context, id string) (err error) {
	request := RemoveCompleteRequest{Id: id}
	response, err := e.RemoveCompleteEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(RemoveCompleteResponse).Err
}

// Delete implements Service. Primarily useful in a client.
func (e Endpoints) Delete(ctx context.Context, id string) (err error) {
	request := DeleteRequest{Id: id}
	response, err := e.DeleteEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteResponse).Err
}
