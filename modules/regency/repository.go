package regency

import "github.com/jmoiron/sqlx"

type Repository interface {
	FindByProvinceCode(provinceCode string) ([]Regency, error)
}

type repository struct {
	db *sqlx.DB
}

func newRepository(db *sqlx.DB) repository {
	return repository{
		db: db,
	}
}

func (r repository) FindByProvinceCode(provinceCode string) ([]Regency, error) {
	var regencies []Regency
	if err := r.db.Select(&regencies, "SELECT * FROM regencies WHERE province_code = ?", provinceCode); err != nil {
		return nil, err
	}
	return regencies, nil
}
