package subdistrict

type Subdistrict struct {
	Code         string `db:"code"`
	Name         string `db:"name"`
	DistrictCode string `db:"district_code"`
	Postcode     string `db:"postcode"`
}

type SubdistrictSearch struct {
	Code         string `db:"code"`
	Name         string `db:"name"`
	DistrictCode string `db:"district_code"`
	Postcode     string `db:"postcode"`
	DistrictName string `db:"district_name"`
	RegencyCode  string `db:"regency_code"`
	RegencyName  string `db:"regency_name"`
	ProvinceCode string `db:"province_code"`
	ProvinceName string `db:"province_name"`
	Address      string `db:"address"`
}
