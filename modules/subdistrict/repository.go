package subdistrict

import "github.com/jmoiron/sqlx"

type Repository interface {
	FindByDistrictCode(districtCode string) ([]Subdistrict, error)
}

type repository struct {
	db *sqlx.DB
}

func newRepository(db *sqlx.DB) repository {
	return repository{
		db: db,
	}
}

func (r repository) FindByDistrictCode(districtCode string) ([]Subdistrict, error) {
	var subdistricts []Subdistrict
	if err := r.db.Select(&subdistricts, "SELECT * FROM subdistricts WHERE district_code = ?", districtCode); err != nil {
		return nil, err
	}
	return subdistricts, nil
}
