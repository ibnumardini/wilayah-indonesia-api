package district

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

func (h handler) FindByRegencyCode(w http.ResponseWriter, r *http.Request) {
	regencyCode := chi.URLParam(r, "regencyCode")

	districts, err := h.service.FindByRegencyCode(regencyCode)
	if err != nil {
		helper.ResponseError(w, helper.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	if len(districts) == 0 {
		helper.ResponseError(w, helper.Response{Status: http.StatusNotFound})
		return
	}

	helper.ResponseSuccess(w, helper.Response{
		Result: districts,
	})
}
