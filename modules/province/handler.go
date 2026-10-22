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

// FindAll godoc
// @Summary Get all provinces
// @Description Retrieves all provinces in Indonesia.
// @Tags province
// @Produce json
// @Success 200 {object} helper.Response
// @Failure 500 {object} helper.Response
// @Router /provinces [get]
func (h handler) FindAll(w http.ResponseWriter, r *http.Request) {
	provinces, err := h.service.FindAll(r.Context())
	if err != nil {
		helper.ResponseError(w, helper.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	responses := make([]Response, 0, len(provinces))
	for _, province := range provinces {
		responses = append(responses, Response(province))
	}

	helper.ResponseSuccess(w, helper.Response{Result: responses})
}

// FindByCode godoc
// @Summary Get province by code
// @Description Retrieves a province by its code.
// @Tags province
// @Produce json
// @Param code path string true "Province code" example(11)
// @Success 200 {object} helper.Response
// @Failure 404 {object} helper.Response
// @Failure 500 {object} helper.Response
// @Router /provinces/{code} [get]
func (h handler) FindByCode(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")

	province, err := h.service.FindByCode(r.Context(), code)
	if err != nil {
		helper.ResponseError(w, helper.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	if province == nil {
		helper.ResponseError(w, helper.Response{Status: http.StatusNotFound})
		return
	}

	helper.ResponseSuccess(w, helper.Response{Result: Response(*province)})
}
