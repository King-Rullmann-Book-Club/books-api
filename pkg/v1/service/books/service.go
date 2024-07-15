package books

import "context"

type Service interface {
    GetBook(ctx context.Context, id string) (Book, error)
}

type Book struct {
    Name string
}

type bookSvc struct {}

func NewService() Service {
    return &bookSvc{}
}
