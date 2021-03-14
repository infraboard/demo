package impl

import (
	"context"

	"github.com/infraboard/demo/conf"
	"github.com/infraboard/demo/pkg/example"
)

func (s *service) CreateBook(ctx context.Context, req *example.CreateBookRequest) (*example.Book, error) {
	tk, err := conf.GetTokenFromGrpcInCtx(ctx)
	if err != nil {
		return nil, err
	}
	s.log.Debug(tk)
	return example.NewBook(req), nil
}

func (s *service) QueryBook(ctx context.Context, req *example.QueryBookRequest) (*example.BookSet, error) {
	return example.NewBookSet(), nil
}
