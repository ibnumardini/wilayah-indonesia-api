package district

import "context"

type Repository interface {
	FindByRegencyCode(ctx context.Context, regencyCode string) ([]District, error)
}
