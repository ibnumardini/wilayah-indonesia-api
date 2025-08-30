package regency

import "github.com/go-chi/chi"

func newRouter(h *handler) *chi.Mux {
	r := chi.NewMux()

	r.Get("/{provinceCode}", h.FindByProvinceCode)

	return r
}
