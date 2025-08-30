package regency

type Regency struct {
	Code         string `db:"code"`
	Name         string `db:"name"`
	ProvinceCode string `db:"province_code"`
}
