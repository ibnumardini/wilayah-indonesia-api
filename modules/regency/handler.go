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

func (h handler) FindByProvinceCode(w http.ResponseWriter, r *http.Request) {
	provinceCode := chi.URLParam(r, "provinceCode")

	provinces, err := h.service.FindByProvinceCode(provinceCode)
	if err != nil {
		helper.ResponseError(w, helper.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	helper.ResponseSuccess(w, helper.Response{
		Result: provinces,
	})
}
