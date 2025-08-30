package district

type District struct {
	Code        string `db:"code"`
	Name        string `db:"name"`
	RegencyCode string `db:"regency_code"`
}
