package province

import "context"

type Repository interface {
	FindAll(ctx context.Context) ([]Province, error)
	FindByCode(ctx context.Context, code string) (*Province, error)
}
