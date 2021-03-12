package impl

import (
	"context"

	"github.com/infraboard/demo/pkg/example"
)

func (s *service) CreateBook(ctx context.Context, req *example.CreateBookRequest) (*example.Book, error) {
	return &example.Book{}, nil
}
