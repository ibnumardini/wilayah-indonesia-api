package district

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

// FindByRegencyCode godoc
// @Summary Get districts by regency code
// @Description Retrieves all districts within a regency.
// @Tags district
// @Produce json
// @Param regencyCode path string true "Regency code" example(1101)
// @Success 200 {object} httpx.Response
// @Failure 404 {object} httpx.Response
// @Failure 500 {object} httpx.Response
// @Router /districts/{regencyCode} [get]
func (h handler) FindByRegencyCode(w http.ResponseWriter, r *http.Request) {
	regencyCode := chi.URLParam(r, "regencyCode")

	districts, err := h.service.FindByRegencyCode(r.Context(), regencyCode)
	if err != nil {
		httpx.InternalError(w, "district.find_by_regency_code", err)
		return
	}

	if len(districts) == 0 {
		httpx.ResponseError(w, httpx.Response{Status: http.StatusNotFound})
		return
	}

	responses := make([]Response, 0, len(districts))
	for _, district := range districts {
		responses = append(responses, Response{
			Code: district.Code,
			Name: district.Name,
		})
	}

	httpx.ResponseSuccess(w, httpx.Response{Result: responses})
}
