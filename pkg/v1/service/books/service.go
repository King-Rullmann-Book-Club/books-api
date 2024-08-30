package books

import (
	"context"

	"github.com/King-Rullmann-Book-Club/books-api/pkg/v1/service/storage"
)

type Service interface {
	GetBook(ctx context.Context, id string) (*Book, error)
}

type Book struct {
	Id    uint   `json:"id"`
	Title string `json:"title"`
}

type bookSvc struct {
    db storage.Transactor
}

func NewService(db storage.Transactor) Service {
	return &bookSvc{db}
}
