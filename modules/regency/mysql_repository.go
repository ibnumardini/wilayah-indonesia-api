package regency

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

func (r mysqlRepository) FindByProvinceCode(ctx context.Context, provinceCode string) ([]Regency, error) {
	var regencies []Regency
	if err := r.db.SelectContext(ctx, &regencies, "SELECT * FROM regencies WHERE province_code = ?", provinceCode); err != nil {
		return nil, err
	}
	return regencies, nil
}
