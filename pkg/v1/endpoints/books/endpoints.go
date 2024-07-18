package books

import (
    "context"
    "github.com/King-Rullmann-Book-Club/books-api/pkg/v1/service/books"
    "github.com/go-kit/kit/endpoint"
)

// COPIED FROM GO-KIT EXAMPLES:
// Endpoints collects all of the endpoints that compose a profile service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
//
// In a server, it's useful for functions that need to operate on a per-endpoint
// basis. For example, you might pass an Endpoints to a function that produces
// an http.Handler, with each method (endpoint) wired up to a specific path. (It
// is probably a mistake in design to invoke the Service methods on the
// Endpoints struct in a server.)
//
// In a client, it's useful to collect individually constructed endpoints into a
// single type that implements the Service interface. For example, you might
// construct individual endpoints using transport/http.NewClient, combine them
// into an Endpoints, and return it to the caller as a Service.
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
