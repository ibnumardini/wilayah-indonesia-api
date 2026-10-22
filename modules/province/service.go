package province

import "context"

type Service interface {
	FindAll(ctx context.Context) ([]Province, error)
	FindByCode(ctx context.Context, code string) (*Province, error)
}

type service struct {
	repo Repository
}

func newService(repo Repository) service {
	return service{
		repo: repo,
	}
}

func (s service) FindAll(ctx context.Context) ([]Province, error) {
	provinces, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return provinces, nil
}

func (s service) FindByCode(ctx context.Context, code string) (*Province, error) {
	province, err := s.repo.FindByCode(ctx, code)
	if err != nil {
		return nil, err
	}

	return province, nil
}
