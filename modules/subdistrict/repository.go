package subdistrict

import "context"

type Repository interface {
	FindByDistrictCode(ctx context.Context, districtCode string) ([]Subdistrict, error)
	FindByQuery(ctx context.Context, searchQuery string, pagination QueryOptions) ([]SubdistrictSearch, error)
	CountByQuery(ctx context.Context, searchQuery string) (int, error)
}
