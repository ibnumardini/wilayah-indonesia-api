package regency

import "context"

type Service interface {
	FindByProvinceCode(ctx context.Context, provinceCode string) ([]Regency, error)
}

type service struct {
	repo Repository
}

func newService(repo Repository) service {
	return service{
		repo: repo,
	}
}

func (s service) FindByProvinceCode(ctx context.Context, provinceCode string) ([]Regency, error) {
	regencies, err := s.repo.FindByProvinceCode(ctx, provinceCode)
	if err != nil {
		return nil, err
	}

	return regencies, nil
}
