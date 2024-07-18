package books

import "context"

type Service interface {
    GetBook(ctx context.Context, id string) (Book, error)
}

type Book struct {
    Title string `json:"title"`
}

type bookSvc struct {}

func NewService() Service {
    return &bookSvc{}
}
