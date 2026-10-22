package regency

import "context"

type Repository interface {
	FindByProvinceCode(ctx context.Context, provinceCode string) ([]Regency, error)
}
