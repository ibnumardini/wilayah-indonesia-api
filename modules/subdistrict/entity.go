package subdistrict

type Subdistrict struct {
	Code         string `db:"code"`
	Name         string `db:"name"`
	DistrictCode string `db:"district_code"`
	Postcode     string `db:"postcode"`
}
