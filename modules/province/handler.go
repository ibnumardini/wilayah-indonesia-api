package province

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

func (h handler) FindAll(w http.ResponseWriter, r *http.Request) {
	provinces, err := h.service.FindAll()
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

func (h handler) FindByCode(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")

	province, err := h.service.FindByCode(code)
	if err != nil {
		helper.ResponseError(w, helper.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	if province == (Response{}) {
		helper.ResponseError(w, helper.Response{Status: http.StatusNotFound})
		return
	}

	helper.ResponseSuccess(w, helper.Response{
		Result: province,
	})
}
