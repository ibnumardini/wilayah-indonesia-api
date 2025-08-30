package subdistrict

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/gorilla/schema"
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

var decoder = schema.NewDecoder()

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

func (h handler) FindByQuery(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		helper.ResponseError(w, helper.Response{Status: http.StatusBadRequest})
		return
	}

	var searchQuery string = strings.ToLower(chi.URLParam(r, "searchQuery"))

	var pagination helper.PaginationRequest
	if err := decoder.Decode(&pagination, r.URL.Query()); err != nil {
		helper.ResponseError(w, helper.Response{Status: http.StatusBadRequest})
		return
	}

	subdistricts, count, err := h.service.FindByQuery(searchQuery, pagination)
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

	paginationResponse := helper.PaginationResponse{
		Total:   count,
		PerPage: pagination.Limit,
		Page:    pagination.Page,
	}

	helper.ResponseSuccess(w, helper.Response{
		Meta:   paginationResponse.ToMeta(),
		Result: subdistricts,
	})
}
