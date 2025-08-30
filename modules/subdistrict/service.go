package subdistrict

type Service interface {
	FindByDistrictCode(districtCode string) ([]Response, error)
}

type service struct {
	repo Repository
}

func newService(repo Repository) service {
	return service{
		repo: repo,
	}
}

func (s service) FindByDistrictCode(districtCode string) ([]Response, error) {
	subdistricts, err := s.repo.FindByDistrictCode(districtCode)
	if err != nil {
		return nil, err
	}

	var responses []Response
	for _, subdistrict := range subdistricts {

		responses = append(responses, Response{
			Code: subdistrict.Code,
			Name: subdistrict.Name,
		})
	}

	return responses, nil
}
