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

func TestListBooks_Success(t *testing.T) {
    svc := NewMockSvc()
    ctx := context.Background()
    svc.On("ListBooks", ctx).Return([]Book{{Title: "Perdido Street Station"}, {Title: "The Scar"}}, nil)

    books, err := svc.ListBooks(ctx)
    assert.Equal(t, books[0].Title, "Perdido Street Station")
    assert.Equal(t, books[1].Title, "The Scar")
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

func (m *mockSvc) ListBooks(ctx context.Context) ([]Book, error) {
    args := m.Called(ctx)
    return args.Get(0).([]Book), args.Error(1)
}

