package subdistrict

import (
	"fmt"

	"github.com/ibnumardini/wilayah-indonesia-api/pkg/helper"
)

type Service interface {
	FindByDistrictCode(districtCode string) ([]Response, error)
	FindByQuery(searchQuery string, pagination helper.PaginationRequest) ([]ResponseSearch, int, error)
}

type service struct {
	repo Repository
}

func newService(repo Repository) service {
	return service{
		repo: repo,
	}
}

func (s service) FindByDistrictCode(districtCode string) ([]Response, error) {
	subdistricts, err := s.repo.FindByDistrictCode(districtCode)
	if err != nil {
		return nil, err
	}

	var responses []Response
	for _, subdistrict := range subdistricts {

		responses = append(responses, Response{
			Code:     subdistrict.Code,
			Name:     subdistrict.Name,
			Postcode: subdistrict.Postcode,
		})
	}

	return responses, nil
}

func (s service) FindByQuery(searchQuery string, pagination helper.PaginationRequest) ([]ResponseSearch, int, error) {
	pagination.Offset = (pagination.Page - 1) * pagination.Limit

	subdistricts, err := s.repo.FindByQuery(searchQuery, pagination)
	if err != nil {
		return nil, 0, err
	}

	count, err := s.repo.CountByQuery(searchQuery)
	if err != nil {
		return nil, 0, err
	}

	var responses []ResponseSearch
	for _, subdistrict := range subdistricts {
		responses = append(responses, ResponseSearch{
			Code:         subdistrict.Code,
			Name:         subdistrict.Name,
			Postcode:     subdistrict.Postcode,
			DistrictCode: subdistrict.DistrictCode,
			DistrictName: subdistrict.DistrictName,
			RegencyCode:  subdistrict.RegencyCode,
			RegencyName:  subdistrict.RegencyName,
			ProvinceCode: subdistrict.ProvinceCode,
			ProvinceName: subdistrict.ProvinceName,
			Address:      fmt.Sprintf("%s, %s, %s, Prov. %s %s", subdistrict.Name, subdistrict.DistrictName, subdistrict.RegencyName, subdistrict.ProvinceName, subdistrict.Postcode),
		})
	}

	return responses, count, nil
}
