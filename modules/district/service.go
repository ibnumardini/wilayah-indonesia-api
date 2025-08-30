package district

type Service interface {
	FindByRegencyCode(regencyCode string) ([]Response, error)
}

type service struct {
	repo Repository
}

func newService(repo Repository) service {
	return service{
		repo: repo,
	}
}

func (s service) FindByRegencyCode(regencyCode string) ([]Response, error) {
	districts, err := s.repo.FindByRegencyCode(regencyCode)
	if err != nil {
		return nil, err
	}

	var responses []Response
	for _, district := range districts {

		responses = append(responses, Response{
			Code: district.Code,
			Name: district.Name,
		})
	}

	return responses, nil
}
