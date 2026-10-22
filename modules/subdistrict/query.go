package subdistrict

import "slices"

type QueryOptions struct {
	Page   int
	Limit  int
	Offset int
	Sort   string
}

func (q *QueryOptions) CheckSort() {
	if !slices.Contains([]string{"asc", "desc"}, q.Sort) {
		q.Sort = "asc"
	}
}
