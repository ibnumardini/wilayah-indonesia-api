package regency

type Service interface {
	FindByProvinceCode(provinceCode string) ([]Response, error)
}

type service struct {
	repo Repository
}

func newService(repo Repository) service {
	return service{
		repo: repo,
	}
}

func (s service) FindByProvinceCode(provinceCode string) ([]Response, error) {
	regencies, err := s.repo.FindByProvinceCode(provinceCode)
	if err != nil {
		return nil, err
	}

	var responses []Response
	for _, regency := range regencies {

		responses = append(responses, Response{
			Code: regency.Code,
			Name: regency.Name,
		})
	}

	return responses, nil
}
