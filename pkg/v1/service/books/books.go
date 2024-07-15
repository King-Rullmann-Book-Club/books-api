package books

import (
    "context"
)

// GetBook returns the details of a given book ID. 
func (s *bookSvc) GetBook(ctx context.Context, id string) (Book, error) {
    return Book{Name: "Perdido Street Station"}, nil
}
