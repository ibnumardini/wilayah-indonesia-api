package province

type Service interface {
	FindAll() ([]Province, error)
}

type service struct {
	repo Repository
}

func newService(repo Repository) service {
	return service{
		repo: repo,
	}
}

func (s service) FindAll() ([]Province, error) {
	return s.repo.FindAll()
}
