package subdistrict

import (
	"fmt"
	"strconv"

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
			RegencyName:  fmt.Sprintf("%s %s", getRegencyType(subdistrict.RegencyCode), subdistrict.RegencyName),
			ProvinceCode: subdistrict.ProvinceCode,
			ProvinceName: subdistrict.ProvinceName,
			Address:      addressBuilder(subdistrict),
		})
	}

	return responses, count, nil
}

func getRegencyType(regencyCode string) string {
	typeCode, _ := strconv.Atoi((regencyCode[len(regencyCode)-2:]))
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
