package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/ibnumardini/wilayah-indonesia-api/modules/district"
	"github.com/ibnumardini/wilayah-indonesia-api/modules/province"
	"github.com/ibnumardini/wilayah-indonesia-api/modules/regency"
	"github.com/ibnumardini/wilayah-indonesia-api/modules/subdistrict"
	"github.com/ibnumardini/wilayah-indonesia-api/pkg/config"
	"github.com/ibnumardini/wilayah-indonesia-api/pkg/db"
	"github.com/ibnumardini/wilayah-indonesia-api/pkg/helper"
	httpSwagger "github.com/swaggo/http-swagger"
)

func init() {
	config.InitConfig()
}

func Handler(w http.ResponseWriter, req *http.Request) {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
	}))

	dbConn, err := db.OpenDbConnection()
	if err != nil {
		helper.ResponseError(w, helper.Response{
			Status:  http.StatusInternalServerError,
			Message: "Database connection failed",
		})
		return
	}

	r.Get("/", welcomeHandler)
	r.Get("/docs/swagger.json", swaggerHandler)
	r.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL("/docs/swagger.json")))
	r.Mount("/provinces", province.LoadModule(dbConn))
	r.Mount("/regencies", regency.LoadModule(dbConn))
	r.Mount("/districts", district.LoadModule(dbConn))
	r.Mount("/subdistricts", subdistrict.LoadModule(dbConn))

	r.ServeHTTP(w, req)
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	helper.ResponseSuccess(w, helper.Response{
		Status:  http.StatusOK,
		Message: "Welcome to Wilayah Indonesia API",
		Result:  struct{}{},
	})
}

func swaggerHandler(w http.ResponseWriter, r *http.Request) {
	httpSwagger.Handler(httpSwagger.URL("/docs/swagger.json"))(w, r)
}
