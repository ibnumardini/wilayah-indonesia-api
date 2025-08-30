package subdistrict

import "github.com/go-chi/chi"

func newRouter(h *handler) *chi.Mux {
	r := chi.NewMux()

	r.Get("/{districtCode}", h.FindByDistrictCode)

	return r
}
