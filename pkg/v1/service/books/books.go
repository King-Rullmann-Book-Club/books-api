package books

import (
	"context"
	"fmt"
	"strconv"
)

// GetBook returns the details of a given book ID.
func (s *bookSvc) GetBook(ctx context.Context, id string) (*Book, error) {
	rid, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		fmt.Errorf("finding %v, %v", id, err)
		return nil, err
	}

    var returnedId uint
    var title string
    if err := s.db.GetRecordById("books", uint(rid), []string{"id", "title"}, &returnedId, &title); err != nil {
		fmt.Errorf("error %v, %v", rid, err)
		return nil, err
    }

	return &Book{returnedId, title}, nil
}

// ListBooks returns all the books in the database (v1)
func (s *bookSvc) ListBooks(ctx context.Context) ([]Book, error) {
    books, err := s.db.GetRecords("books", []string{"id", "title"})
    if err != nil {
        return nil, err
    } 

    var bookList []Book
    for books.Next() {
        var id uint
        var title string
        if err := books.Scan(&id, &title); err != nil {
            return nil, err
        }
        bookList = append(bookList, Book{id, title})
    }
   
    return bookList, nil
}
