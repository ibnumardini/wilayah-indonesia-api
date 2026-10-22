package district

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	FindByRegencyCode(ctx context.Context, regencyCode string) ([]District, error)
}

type repository struct {
	db *sqlx.DB
}

func newRepository(db *sqlx.DB) repository {
	return repository{
		db: db,
	}
}

func (r repository) FindByRegencyCode(ctx context.Context, regencyCode string) ([]District, error) {
	var districts []District
	if err := r.db.SelectContext(ctx, &districts, "SELECT * FROM districts WHERE regency_code = ?", regencyCode); err != nil {
		return nil, err
	}
	return districts, nil
}
