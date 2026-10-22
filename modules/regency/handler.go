package regency

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/ibnumardini/wilayah-indonesia-api/internal/httpx"
)

type handler struct {
	service Service
}

func newHandler(service Service) handler {
	return handler{
		service: service,
	}
}

// FindByProvinceCode godoc
// @Summary Get regencies by province code
// @Description Retrieves all regencies within a province.
// @Tags regencies
// @Produce json
// @Param provinceCode path string true "Province code" example(11)
// @Success 200 {object} httpx.Response
// @Failure 404 {object} httpx.Response
// @Failure 500 {object} httpx.Response
// @Router /regencies/{provinceCode} [get]
func (h handler) FindByProvinceCode(w http.ResponseWriter, r *http.Request) {
	provinceCode := chi.URLParam(r, "provinceCode")

	regencies, err := h.service.FindByProvinceCode(r.Context(), provinceCode)
	if err != nil {
		httpx.InternalError(w, "regency.find_by_province_code", err)
		return
	}

	if len(regencies) == 0 {
		httpx.ResponseError(w, httpx.Response{Status: http.StatusNotFound})
		return
	}

	responses := make([]Response, 0, len(regencies))
	for _, regency := range regencies {
		responses = append(responses, Response{
			Code: regency.Code,
			Name: regency.Name,
		})
	}

	httpx.ResponseSuccess(w, httpx.Response{Result: responses})
}
