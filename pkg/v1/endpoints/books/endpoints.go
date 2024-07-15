package books

import (
    "context"
    "github.com/King-Rullmann-Book-Club/books-api/pkg/v1/service/books"
    "github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
    GetBook endpoint.Endpoint
}

func MakeEndpoints(s books.Service) *Endpoints {
    return &Endpoints{
        GetBook: makeGetBookEndpoint(s),
    }
}

func makeGetBookEndpoint(s books.Service) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetBookRequest)
		b, err := s.GetBook(ctx, req.ID)
        return GetBookResponse{Book: b, Err: err}, nil
	}
}
