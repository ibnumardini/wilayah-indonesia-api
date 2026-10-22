package district

import "context"

type Service interface {
	FindByRegencyCode(ctx context.Context, regencyCode string) ([]District, error)
}

type service struct {
	repo Repository
}

func newService(repo Repository) service {
	return service{
		repo: repo,
	}
}

func (s service) FindByRegencyCode(ctx context.Context, regencyCode string) ([]District, error) {
	districts, err := s.repo.FindByRegencyCode(ctx, regencyCode)
	if err != nil {
		return nil, err
	}

	return districts, nil
}
