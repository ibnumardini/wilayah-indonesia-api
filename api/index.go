package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/ibnumardini/wilayah-indonesia-api/modules/province"
	"github.com/ibnumardini/wilayah-indonesia-api/modules/regency"
	"github.com/ibnumardini/wilayah-indonesia-api/pkg/config"
	"github.com/ibnumardini/wilayah-indonesia-api/pkg/db"
	"github.com/ibnumardini/wilayah-indonesia-api/pkg/helper"
)

func init() {
	config.InitConfig()
}

func Handler(w http.ResponseWriter, req *http.Request) {
	r := chi.NewRouter()

	dbConn, err := db.OpenDbConnection()
	if err != nil {
		helper.ResponseError(w, helper.Response{
			Status:  http.StatusInternalServerError,
			Message: "Database connection failed",
		})
		return
	}

	r.Mount("/provinces", province.LoadModule(dbConn))
	r.Mount("/regencies", regency.LoadModule(dbConn))

	r.ServeHTTP(w, req)
}
