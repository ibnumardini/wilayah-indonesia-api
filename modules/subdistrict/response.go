package subdistrict

type Response struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	Postcode string `json:"postcode"`
}

type ResponseSearch struct {
	Code         string `json:"code"`
	Name         string `json:"name"`
	DistrictCode string `json:"district_code"`
	DistrictName string `json:"district_name"`
	RegencyCode  string `json:"regency_code"`
	RegencyName  string `json:"regency_name"`
	ProvinceCode string `json:"province_code"`
	ProvinceName string `json:"province_name"`
	Postcode     string `json:"postcode"`
	Address      string `json:"address"`
}
