// @title Wilayah Indonesia API
// @version 1.0
// @description REST API for Indonesian administrative regions.
// @BasePath /
// @schemes http https
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/ibnumardini/wilayah-indonesia-api/docs"
	"github.com/ibnumardini/wilayah-indonesia-api/modules/district"
	"github.com/ibnumardini/wilayah-indonesia-api/modules/province"
	"github.com/ibnumardini/wilayah-indonesia-api/modules/regency"
	"github.com/ibnumardini/wilayah-indonesia-api/modules/subdistrict"
	"github.com/ibnumardini/wilayah-indonesia-api/pkg/config"
	"github.com/ibnumardini/wilayah-indonesia-api/pkg/db"
	"github.com/ibnumardini/wilayah-indonesia-api/pkg/helper"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	if err := config.InitConfig(); err != nil {
		log.Fatalf("initialize config: %v", err)
	}

	docs.SwaggerInfo.Host = config.C.Server.SwaggerHost

	dbConn, err := db.OpenDbConnection()
	if err != nil {
		log.Fatalf("connect to database: %v", err)
	}
	defer dbConn.Close()

	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*"},
	}))

	r.Get("/", welcomeHandler)
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
		httpSwagger.UIConfig(map[string]string{
			"tagsSorter": "function(a, b) { var order = ['welcome', 'province', 'regencies', 'district', 'subdistrict']; return order.indexOf(a) - order.indexOf(b); }",
		}),
	))

	r.Mount("/provinces", province.LoadModule(dbConn))
	r.Mount("/regencies", regency.LoadModule(dbConn))
	r.Mount("/districts", district.LoadModule(dbConn))
	r.Mount("/subdistricts", subdistrict.LoadModule(dbConn))

	addr := fmt.Sprintf(":%d", config.C.Server.Port)
	log.Printf("starting %s on %s", config.C.Server.AppName, addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("start server: %v", err)
	}
}

// Welcome godoc
// @Summary Welcome message
// @Description Returns a welcome message for the API.
// @Tags welcome
// @Produce json
// @Success 200 {object} helper.Response
// @Router / [get]
func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	helper.ResponseSuccess(w, helper.Response{
		Status:  http.StatusOK,
		Message: "Welcome to Wilayah Indonesia API",
		Result:  struct{}{},
	})
}
