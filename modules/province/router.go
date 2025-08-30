package province

import "github.com/go-chi/chi"

func newRouter(h *handler) *chi.Mux {
	r := chi.NewMux()

	r.Get("/", h.FindAll)
	r.Get("/{code}", h.FindByCode)

	return r
}
