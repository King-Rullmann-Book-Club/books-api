package books

import (
	"context"
	"fmt"
	"strconv"

	"github.com/King-Rullmann-Book-Club/books-api/pkg/v1/service/storage"
)

// GetBook returns the details of a given book ID.
func (s *bookSvc) GetBook(ctx context.Context, id string) (Book, error) {
	rid, err := strconv.ParseUint(id, 10, 0)

	if err != nil {
		fmt.Errorf("finding %v, %v", id, err)
		return Book{}, err
	}
	return getBookById(uint(rid))
}

// GetBook returns the details of a given book ID.
func getBookById(rid uint) (Book, error) {
	var id uint
	var title string
	t := storage.NewTransactor()
	err := t.GetRecordById("books", rid, []string{"id", "title"}, &id, &title)

	if err != nil {
		fmt.Errorf("error %v, %v", rid, err)
		return Book{}, err
	}
	return Book{Id: id, Title: title}, nil
}
