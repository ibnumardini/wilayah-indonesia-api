package province

import (
	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
)

func LoadModule(dbConn *sqlx.DB) *chi.Mux {
	repository := newMySQLRepository(dbConn)
	service := newService(&repository)
	handler := newHandler(&service)

	return newRouter(&handler)
}
