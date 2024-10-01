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

