package province

import (
	"net/http"

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
