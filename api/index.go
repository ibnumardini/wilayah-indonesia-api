package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/ibnumardini/wilayah-indonesia-api/handlers"
)

func Handler(w http.ResponseWriter, req *http.Request) {
	r := chi.NewRouter()

	r.Get("/", handlers.NewHelloWorldHandler().ServeHTTP)

	r.ServeHTTP(w, req)
}
