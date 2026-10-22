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

// FindByDistrictCode godoc
// @Summary Get subdistricts by district code
// @Description Retrieves all subdistricts within a district.
// @Tags subdistrict
// @Produce json
// @Param districtCode path string true "District code" example(1101010)
// @Success 200 {object} helper.Response
// @Failure 404 {object} helper.Response
// @Failure 500 {object} helper.Response
// @Router /subdistricts/{districtCode} [get]
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

// FindByQuery godoc
// @Summary Search subdistricts
// @Description Searches subdistricts by name or other criteria with pagination.
// @Tags subdistrict
// @Produce json
// @Param searchQuery path string true "Search query" example(jakarta)
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of results per page" default(10)
// @Param sort query string false "Sort order" Enums(asc,desc) default(asc)
// @Success 200 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Failure 404 {object} helper.Response
// @Failure 500 {object} helper.Response
// @Router /subdistricts/maps/{searchQuery} [get]
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

	pagination.CheckSort()

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
