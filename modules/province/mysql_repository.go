package province

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

func (r mysqlRepository) FindAll(ctx context.Context) ([]Province, error) {
	var provinces []Province
	if err := r.db.SelectContext(ctx, &provinces, "SELECT * FROM provinces"); err != nil {
		return nil, err
	}
	return provinces, nil
}

func (r mysqlRepository) FindByCode(ctx context.Context, code string) (*Province, error) {
	var provinces []Province
	if err := r.db.SelectContext(ctx, &provinces, "SELECT * FROM provinces WHERE code = ?", code); err != nil {
		return nil, err
	}
	if len(provinces) == 0 {
		return nil, nil
	}
	return &provinces[0], nil
}
