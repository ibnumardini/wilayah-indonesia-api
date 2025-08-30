package subdistrict

import (
	"github.com/ibnumardini/wilayah-indonesia-api/pkg/helper"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	FindByDistrictCode(districtCode string) ([]Subdistrict, error)
	FindByQuery(searchQuery string, pagination helper.PaginationRequest) ([]SubdistrictSearch, error)
	CountByQuery(searchQuery string) (int, error)
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

func (r repository) FindByQuery(searchQuery string, pagination helper.PaginationRequest) ([]SubdistrictSearch, error) {
	var subdistricts []SubdistrictSearch

	var query = `SELECT s.code, s.name, s.district_code, p.postcode, d.name as district_name, r.code as regency_code, r.name as regency_name,
				pr.code as province_code, pr.name as province_name FROM subdistricts as s
				JOIN postcodes p ON s.code = p.subdistrict_code
				JOIN districts d ON s.district_code = d.code
				JOIN regencies r ON d.regency_code = r.code
				JOIN provinces pr ON r.province_code = pr.code
				WHERE s.name LIKE ? OR p.postcode LIKE ? LIMIT ? OFFSET ?`

	if err := r.db.Select(&subdistricts, query, "%"+searchQuery+"%", "%"+searchQuery+"%", pagination.Limit, pagination.Offset); err != nil {
		return nil, err
	}

	return subdistricts, nil
}

func (r repository) CountByQuery(searchQuery string) (int, error) {
	var count int

	var query = `SELECT COUNT(*) FROM subdistricts as s
				JOIN postcodes p ON s.code = p.subdistrict_code WHERE s.name LIKE ? OR p.postcode LIKE ?`

	if err := r.db.Get(&count, query, "%"+searchQuery+"%", "%"+searchQuery+"%"); err != nil {
		return 0, err
	}

	return count, nil
}
