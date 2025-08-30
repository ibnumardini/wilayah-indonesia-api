package district

import "github.com/go-chi/chi"

func newRouter(h *handler) *chi.Mux {
	r := chi.NewMux()

	r.Get("/{regencyCode}", h.FindByRegencyCode)

	return r
}
