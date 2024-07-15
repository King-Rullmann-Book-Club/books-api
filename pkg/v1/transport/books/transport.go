package books

import (
    "bytes"
    "context"
    "encoding/json"
    "errors"
    "io"
    "net/http"

    svc "github.com/King-Rullmann-Book-Club/books-api/pkg/v1/service/books"
    ep "github.com/King-Rullmann-Book-Club/books-api/pkg/v1/endpoints/books"

    "github.com/gorilla/mux"
    httptransport "github.com/go-kit/kit/transport/http"
)

// NewTransport returns a new mux router for exposed endpoints
func NewTransport(s svc.Service) http.Handler {
    r := mux.NewRouter()
    e := ep.MakeEndpoints(s)
    options := []httptransport.ServerOption{}

    // GET /books/:id                 Get a book by ID

    r.Methods(http.MethodGet).Path("/books/{id}").Handler(httptransport.NewServer(
        e.GetBook,
        decodeGetBookRequest,
        encodeResponse,
        options...,
    ))

    return r
}

// decodeGetBookRequest decodes a request with the following format: GET /books/{id}
func decodeGetBookRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
    vars := mux.Vars(r)
    id, ok := vars["id"]
    if !ok {
        return nil, errors.New("Unable to get parameter, bad route")
    }
    return ep.GetBookRequest{ID: id}, nil
}

// errorer is implemented by all concrete response types that may contain
// errors. It allows us to change the HTTP response code without needing to
// trigger an endpoint (transport-level) error. For more information, read the
// big comment in endpoints.go.
type errorer interface {
    error() error
}

// encodeResponse is the common method to encode all response types to the
// client. I chose to do it this way because, since we're using JSON, there's no
// reason to provide anything more specific. It's certainly possible to
// specialize on a per-response (per-method) basis.
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
    if e, ok := response.(errorer); ok && e.error() != nil {
        // Not a Go kit transport error, but a business-logic error.
        // Provide those as HTTP errors.
        encodeError(ctx, e.error(), w)
        return nil
    }
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    return json.NewEncoder(w).Encode(response)
}

// encodeRequest likewise JSON-encodes the request to the HTTP request body.
// Don't use it directly as a transport/http.Client EncodeRequestFunc:
// profilesvc endpoints require mutating the HTTP method and request path.
func encodeRequest(_ context.Context, req *http.Request, request interface{}) error {
    var buf bytes.Buffer
    err := json.NewEncoder(&buf).Encode(request)
    if err != nil {
        return err
    }
    req.Body = io.NopCloser(&buf)
    return nil
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
    if err == nil {
        panic("encodeError with nil error")
    }
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    w.WriteHeader(http.StatusInternalServerError)
    json.NewEncoder(w).Encode(map[string]interface{}{
        "error": err.Error(),
    })
}

