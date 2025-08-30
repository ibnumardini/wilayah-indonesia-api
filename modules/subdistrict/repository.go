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

	var query = `SELECT s.code, s.name, s.district_code, p.postcode FROM subdistricts as s
				JOIN postcodes p ON s.code = p.subdistrict_code WHERE s.district_code = ?`

	if err := r.db.Select(&subdistricts, query, districtCode); err != nil {
		return nil, err
	}

	return subdistricts, nil
}
