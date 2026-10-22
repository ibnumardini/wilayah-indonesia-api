package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/ibnumardini/wilayah-indonesia-api/modules/district"
	"github.com/ibnumardini/wilayah-indonesia-api/modules/province"
	"github.com/ibnumardini/wilayah-indonesia-api/modules/regency"
	"github.com/ibnumardini/wilayah-indonesia-api/modules/subdistrict"
	"github.com/ibnumardini/wilayah-indonesia-api/pkg/config"
	"github.com/ibnumardini/wilayah-indonesia-api/pkg/db"
	"github.com/ibnumardini/wilayah-indonesia-api/pkg/env"
	"github.com/ibnumardini/wilayah-indonesia-api/pkg/helper"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	if err := config.InitConfig(); err != nil {
		log.Fatalf("initialize config: %v", err)
	}

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
	r.Get("/docs/swagger.json", swaggerHandler)
	r.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL("/docs/swagger.json")))

	r.Mount("/provinces", province.LoadModule(dbConn))
	r.Mount("/regencies", regency.LoadModule(dbConn))
	r.Mount("/districts", district.LoadModule(dbConn))
	r.Mount("/subdistricts", subdistrict.LoadModule(dbConn))

	addr := ":" + env.Get("PORT", "8080")
	log.Printf("starting %s on %s", config.C.Server.AppName, addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("start server: %v", err)
	}
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
