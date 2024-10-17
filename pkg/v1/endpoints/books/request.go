package books

import (
	"github.com/King-Rullmann-Book-Club/books-api/pkg/v1/service/books"
)

type GetBookRequest struct {
	ID string
}

type GetBookResponse struct {
	Book books.Book `json:"book,omitempty"`
	Err  error      `json:"err,omitempty"`
}

type ListBooksRequest struct{}

type ListBooksResponse struct {
    Books []books.Book `json:"books,omitempty"`
    Err   error        `json:"err,omitempty"`
}
