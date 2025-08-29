package province

type Service interface {
	FindAll() ([]Response, error)
}

type service struct {
	repo Repository
}

func newService(repo Repository) service {
	return service{
		repo: repo,
	}
}

func (s service) FindAll() ([]Response, error) {
	provinces, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	var responses []Response
	for _, province := range provinces {
		responses = append(responses, Response(province))
	}

	return responses, nil
}
