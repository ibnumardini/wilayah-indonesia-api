package subdistrict

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

func (h handler) FindByDistrictCode(w http.ResponseWriter, r *http.Request) {
	districtCode := chi.URLParam(r, "districtCode")

	subdistricts, err := h.service.FindByDistrictCode(districtCode)
	if err != nil {
		helper.ResponseError(w, helper.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	if len(subdistricts) == 0 {
		helper.ResponseError(w, helper.Response{Status: http.StatusNotFound})
		return
	}

	helper.ResponseSuccess(w, helper.Response{
		Result: subdistricts,
	})
}
