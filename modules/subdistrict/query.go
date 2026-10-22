package subdistrict

import (
	"fmt"
	"strconv"

	"github.com/asaskevich/govalidator"
)

type QueryOptions struct {
	Page   int
	Limit  int
	Offset int
	Sort   string
}

func (q *QueryOptions) CheckSort() {
	if !govalidator.IsIn(q.Sort, "asc", "desc") {
		q.Sort = "asc"
	}
}

func (q QueryOptions) Validate() error {
	if !govalidator.IsInt(strconv.Itoa(q.Page)) || q.Page < 1 {
		return fmt.Errorf("page must be a positive integer")
	}
	if !govalidator.IsInt(strconv.Itoa(q.Limit)) || q.Limit < 1 || q.Limit > 100 {
		return fmt.Errorf("limit must be an integer between 1 and 100")
	}
	if !govalidator.IsIn(q.Sort, "asc", "desc") {
		return fmt.Errorf("sort must be asc or desc")
	}
	return nil
}
