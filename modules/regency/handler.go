package regency

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/ibnumardini/wilayah-indonesia-api/pkg/helper"
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
// @Success 200 {object} helper.Response
// @Failure 404 {object} helper.Response
// @Failure 500 {object} helper.Response
// @Router /regencies/{provinceCode} [get]
func (h handler) FindByProvinceCode(w http.ResponseWriter, r *http.Request) {
	provinceCode := chi.URLParam(r, "provinceCode")

	regencies, err := h.service.FindByProvinceCode(r.Context(), provinceCode)
	if err != nil {
		helper.ResponseError(w, helper.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	if len(regencies) == 0 {
		helper.ResponseError(w, helper.Response{Status: http.StatusNotFound})
		return
	}

	responses := make([]Response, 0, len(regencies))
	for _, regency := range regencies {
		responses = append(responses, Response{
			Code: regency.Code,
			Name: regency.Name,
		})
	}

	helper.ResponseSuccess(w, helper.Response{Result: responses})
}
