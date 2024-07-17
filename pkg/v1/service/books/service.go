package books

import "context"

type Service interface {
    GetBook(ctx context.Context, id string) (Book, error)
}

type Book struct {
	Id    uint   `json:"id"`
    Title string `json:"title"`
}

type bookSvc struct {}

func NewService() Service {
    return &bookSvc{}
}
