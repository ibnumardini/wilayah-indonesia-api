package province

import (
	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
)

func LoadModule(dbConn *sqlx.DB) *chi.Mux {
	repo := newRepository(dbConn)
	service := newService(&repo)
	handler := newHandler(&service)

	r := newRouter(&handler)

	return r
}
