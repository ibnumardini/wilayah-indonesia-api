package subdistrict

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type mysqlRepository struct {
	db *sqlx.DB
}

func newMySQLRepository(db *sqlx.DB) mysqlRepository {
	return mysqlRepository{db: db}
}

func (r mysqlRepository) FindByDistrictCode(ctx context.Context, districtCode string) ([]Subdistrict, error) {
	var subdistricts []Subdistrict
	query := `SELECT s.code, s.name, s.district_code, p.postcode FROM subdistricts as s
				JOIN postcodes p ON s.code = p.subdistrict_code WHERE s.district_code = ?`

	if err := r.db.SelectContext(ctx, &subdistricts, query, districtCode); err != nil {
		return nil, err
	}
	return subdistricts, nil
}

func (r mysqlRepository) FindByQuery(ctx context.Context, searchQuery string, pagination QueryOptions) ([]SubdistrictSearch, error) {
	var subdistricts []SubdistrictSearch
	query := `SELECT s.code, s.name, s.district_code, p.postcode, d.name as district_name, r.code as regency_code, r.name as regency_name,
				pr.code as province_code, pr.name as province_name FROM subdistricts as s
				JOIN postcodes p ON s.code = p.subdistrict_code
				JOIN districts d ON s.district_code = d.code
				JOIN regencies r ON d.regency_code = r.code
				JOIN provinces pr ON r.province_code = pr.code
				WHERE LOWER(s.name) LIKE ? OR p.postcode LIKE ? ORDER BY s.name ` + pagination.Sort + ` LIMIT ? OFFSET ?`

	if err := r.db.SelectContext(ctx, &subdistricts, query, "%"+searchQuery+"%", "%"+searchQuery+"%", pagination.Limit, pagination.Offset); err != nil {
		return nil, err
	}
	return subdistricts, nil
}

func (r mysqlRepository) CountByQuery(ctx context.Context, searchQuery string) (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM subdistricts as s
				JOIN postcodes p ON s.code = p.subdistrict_code WHERE LOWER(s.name) LIKE ? OR p.postcode LIKE ?`

	if err := r.db.GetContext(ctx, &count, query, "%"+searchQuery+"%", "%"+searchQuery+"%"); err != nil {
		return 0, err
	}
	return count, nil
}
