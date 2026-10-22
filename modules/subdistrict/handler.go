package subdistrict

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
	"github.com/gorilla/schema"
	"github.com/ibnumardini/wilayah-indonesia-api/internal/httpx"
	"github.com/ibnumardini/wilayah-indonesia-api/internal/pagination"
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

type paginationQuery struct {
	Page  int    `schema:"page,default:1"`
	Limit int    `schema:"limit,default:10"`
	Sort  string `schema:"sort,default:asc"`
}

// FindByDistrictCode godoc
// @Summary Get subdistricts by district code
// @Description Retrieves all subdistricts within a district.
// @Tags subdistrict
// @Produce json
// @Param districtCode path string true "District code" example(1101010)
// @Success 200 {object} httpx.Response
// @Failure 404 {object} httpx.Response
// @Failure 500 {object} httpx.Response
// @Router /subdistricts/{districtCode} [get]
func (h handler) FindByDistrictCode(w http.ResponseWriter, r *http.Request) {
	districtCode := chi.URLParam(r, "districtCode")

	subdistricts, err := h.service.FindByDistrictCode(r.Context(), districtCode)
	if err != nil {
		httpx.InternalError(w, "subdistrict.find_by_district_code", err)
		return
	}

	if len(subdistricts) == 0 {
		httpx.ResponseError(w, httpx.Response{Status: http.StatusNotFound})
		return
	}

	responses := make([]Response, 0, len(subdistricts))
	for _, subdistrict := range subdistricts {
		responses = append(responses, Response{
			Code:     subdistrict.Code,
			Name:     subdistrict.Name,
			Postcode: subdistrict.Postcode,
		})
	}

	httpx.ResponseSuccess(w, httpx.Response{Result: responses})
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
// @Success 200 {object} httpx.Response
// @Failure 400 {object} httpx.Response
// @Failure 404 {object} httpx.Response
// @Failure 500 {object} httpx.Response
// @Router /subdistricts/maps/{searchQuery} [get]
func (h handler) FindByQuery(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		httpx.ResponseError(w, httpx.Response{Status: http.StatusBadRequest})
		return
	}

	var searchQuery string = strings.ToLower(chi.URLParam(r, "searchQuery"))

	var paginationParams paginationQuery
	if err := decoder.Decode(&paginationParams, r.URL.Query()); err != nil {
		httpx.ResponseError(w, httpx.Response{Status: http.StatusBadRequest})
		return
	}

	options := QueryOptions{
		Page:  paginationParams.Page,
		Limit: paginationParams.Limit,
		Sort:  paginationParams.Sort,
	}
	options.CheckSort()
	if err := options.Validate(); err != nil {
		httpx.ResponseError(w, httpx.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	subdistricts, count, err := h.service.FindByQuery(r.Context(), searchQuery, options)
	if err != nil {
		httpx.InternalError(w, "subdistrict.find_by_query", err)
		return
	}

	if len(subdistricts) == 0 {
		httpx.ResponseError(w, httpx.Response{Status: http.StatusNotFound})
		return
	}

	paginationMetadata := pagination.Metadata{
		Total:   count,
		PerPage: options.Limit,
		Page:    options.Page,
	}

	responses := make([]ResponseSearch, 0, len(subdistricts))
	for _, subdistrict := range subdistricts {
		responses = append(responses, toResponseSearch(subdistrict))
	}

	httpx.ResponseSuccess(w, httpx.Response{
		Meta:   paginationMetadata.ToMap(),
		Result: responses,
	})
}

func toResponseSearch(subdistrict SubdistrictSearch) ResponseSearch {
	return ResponseSearch{
		Code:         subdistrict.Code,
		Name:         subdistrict.Name,
		Postcode:     subdistrict.Postcode,
		DistrictCode: subdistrict.DistrictCode,
		DistrictName: subdistrict.DistrictName,
		RegencyCode:  subdistrict.RegencyCode,
		RegencyName:  fmt.Sprintf("%s %s", getRegencyType(subdistrict.RegencyCode), subdistrict.RegencyName),
		ProvinceCode: subdistrict.ProvinceCode,
		ProvinceName: subdistrict.ProvinceName,
		Address:      addressBuilder(subdistrict),
	}
}

func getRegencyType(regencyCode string) string {
	typeCode, _ := strconv.Atoi(regencyCode[len(regencyCode)-2:])
	if typeCode < 71 {
		return "Kab."
	}
	return "Kota"
}

func addressBuilder(subdistrict SubdistrictSearch) string {
	return fmt.Sprintf("%s, Kec. %s, %s %s, Prov. %s %s",
		subdistrict.Name,
		subdistrict.DistrictName,
		getRegencyType(subdistrict.RegencyCode),
		subdistrict.RegencyName,
		subdistrict.ProvinceName,
		subdistrict.Postcode,
	)
}
