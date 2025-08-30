package province

import "github.com/jmoiron/sqlx"

type Repository interface {
	FindAll() ([]Province, error)
	FindByCode(code string) (Province, error)
}

type repository struct {
	db *sqlx.DB
}

func newRepository(db *sqlx.DB) repository {
	return repository{
		db: db,
	}
}

func (r repository) FindAll() ([]Province, error) {
	var provinces []Province
	if err := r.db.Select(&provinces, "SELECT * FROM provinces"); err != nil {
		return nil, err
	}
	return provinces, nil
}

func (r repository) FindByCode(code string) (Province, error) {
	var province Province
	if err := r.db.Get(&province, "SELECT * FROM provinces WHERE code = ?", code); err != nil {
		return Province{}, err
	}
	return province, nil
}
