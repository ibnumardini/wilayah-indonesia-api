package subdistrict

import (
	"context"
)

type Service interface {
	FindByDistrictCode(ctx context.Context, districtCode string) ([]Subdistrict, error)
	FindByQuery(ctx context.Context, searchQuery string, pagination QueryOptions) ([]SubdistrictSearch, int, error)
}

type service struct {
	repo Repository
}

func newService(repo Repository) service {
	return service{
		repo: repo,
	}
}

func (s service) FindByDistrictCode(ctx context.Context, districtCode string) ([]Subdistrict, error) {
	subdistricts, err := s.repo.FindByDistrictCode(ctx, districtCode)
	if err != nil {
		return nil, err
	}

	return subdistricts, nil
}

func (s service) FindByQuery(ctx context.Context, searchQuery string, pagination QueryOptions) ([]SubdistrictSearch, int, error) {
	pagination.Offset = (pagination.Page - 1) * pagination.Limit

	subdistricts, err := s.repo.FindByQuery(ctx, searchQuery, pagination)
	if err != nil {
		return nil, 0, err
	}

	count, err := s.repo.CountByQuery(ctx, searchQuery)
	if err != nil {
		return nil, 0, err
	}

	return subdistricts, count, nil
}
