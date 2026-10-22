package province

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	FindAll(ctx context.Context) ([]Province, error)
	FindByCode(ctx context.Context, code string) (*Province, error)
}

type repository struct {
	db *sqlx.DB
}

func newRepository(db *sqlx.DB) repository {
	return repository{
		db: db,
	}
}

func (r repository) FindAll(ctx context.Context) ([]Province, error) {
	var provinces []Province
	if err := r.db.SelectContext(ctx, &provinces, "SELECT * FROM provinces"); err != nil {
		return nil, err
	}
	return provinces, nil
}

func (r repository) FindByCode(ctx context.Context, code string) (*Province, error) {
	var provinces []Province

	if err := r.db.SelectContext(ctx, &provinces, "SELECT * FROM provinces WHERE code = ?", code); err != nil {
		return nil, err
	}

	if len(provinces) == 0 {
		return nil, nil
	}

	return &provinces[0], nil
}
