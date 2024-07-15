package books

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetBooks_Success(t *testing.T) {
    svc := NewMockSvc()
    ctx := context.Background()
    svc.On("GetBooks", ctx, "1").Return(Book{Title: "Perdido Street Station"}, nil)

    book, err := svc.GetBooks(ctx, "1")
    assert.Equal(t, book.Title, "Perdido Street Station")
    assert.Nil(t, err) 
}

type mockSvc struct {
    mock.Mock
}

func NewMockSvc() *mockSvc {
    return &mockSvc{}
}

func (m *mockSvc) GetBooks(ctx context.Context, id string) (Book, error) {
    args := m.Called(ctx, id) 
    return args.Get(0).(Book), args.Error(1)
}
