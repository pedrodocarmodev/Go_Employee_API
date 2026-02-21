package employee

import (
	"context"
)

type Repository interface {
	GetAll(ctx context.Context) ([]Employee, error)
	GetById(ctx context.Context, id int) (*Employee, error)
	RegisterEmployee(ctx context.Context, emp *Employee) (int, error)
	Fire(ctx context.Context, id int) (*Employee, error)
	Employ(ctx context.Context, id int) (*Employee, error)
}