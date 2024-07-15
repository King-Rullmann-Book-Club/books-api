package books

import (
    "context"
    "github.com/King-Rullmann-Book-Club/books-api/pkg/v1/service/books"
    "github.com/go-kit/kit/endpoint"
)

// Endpoints represents available endpoints for the book service. 
type Endpoints struct {
    GetBook endpoint.Endpoint
}

// MakeEndpoints returns a list of available service Endpoints. 
func MakeEndpoints(s books.Service) *Endpoints {
    return &Endpoints{
        GetBook: makeGetBookEndpoint(s),
    }
}

// makeGetBookEndpoint formats the get book response. 
func makeGetBookEndpoint(s books.Service) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(GetBookRequest)
        b, err := s.GetBook(ctx, req.ID)
        return GetBookResponse{Book: b, Err: err}, nil
    }
}
