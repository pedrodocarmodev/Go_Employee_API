package employee

import (
	"context"
)

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) GetAll(ctx context.Context) ([]Employee, error) {
	return s.repo.GetAll(ctx)
}

func (s *Service) GetById(ctx context.Context, id int) (*Employee, error) {
	return s.repo.GetById(ctx, id)
}

func (s *Service) RegisterEmployee(ctx context.Context, emp *Employee) (int, error) {
	return s.repo.RegisterEmployee(ctx, emp)
}

func (s *Service) Fire(ctx context.Context, id int) (*Employee, error) {
	return s.repo.Fire(ctx, id)
}

func (s *Service) Employ(ctx context.Context, id int) (*Employee, error) {
	return s.repo.Employ(ctx, id)
}