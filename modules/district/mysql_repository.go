package district

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

func (r mysqlRepository) FindByRegencyCode(ctx context.Context, regencyCode string) ([]District, error) {
	var districts []District
	if err := r.db.SelectContext(ctx, &districts, "SELECT * FROM districts WHERE regency_code = ?", regencyCode); err != nil {
		return nil, err
	}
	return districts, nil
}
