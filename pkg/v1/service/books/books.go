package books

import (
	"context"
	"log"
	"strconv"
)

// GetBook returns the details of a given book ID.
func (s *bookSvc) GetBook(ctx context.Context, id string) (Book, error) {
	rid, err := strconv.ParseUint(id, 10, 0)

	if err != nil {
		println("finding %v, %v", id, err)
		return Book{}, err
	}

	t := NewTransactor()
	return t.GetBookById(uint(rid))
}

// GetBook returns the details of a given book ID.
func (s *transactor) GetBookById(rid uint) (Book, error) {
	row := s.db.QueryRow(`select id, title from books where id = ? limit 1;`, rid)
	println("finding", rid)

	var id uint
	var title string
	err := row.Scan(&id, &title)
	if err != nil {
		log.Default().Printf("error %v, %v", rid, err)

		return Book{}, err
	}
	return Book{Id: id, Title: title}, nil
}
