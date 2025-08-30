package helper

import "slices"

type PaginationRequest struct {
	Page   int `schema:"page,default:1"`
	Limit  int `schema:"limit,default:10"`
	Offset int
	Sort   string `schema:"sort,default:asc"`
}

func (preq *PaginationRequest) CheckSort() {
	if !slices.Contains([]string{"asc", "desc"}, preq.Sort) {
		preq.Sort = "asc"
	}
}

type PaginationResponse struct {
	Total        int `json:"total"`
	PerPage      int `json:"per_page"`
	PreviousPage int `json:"previous_page"`
	Page         int `json:"page"`
	NextPage     int `json:"next_page"`
}

func (pr PaginationResponse) ToMeta() PaginationResponse {
	var previousPage int
	if pr.Page > 1 {
		previousPage = pr.Page - 1
	}

	var nextPage int
	if pr.Page*pr.PerPage < pr.Total {
		nextPage = pr.Page + 1
	}

	return PaginationResponse{
		Total:        pr.Total,
		PerPage:      pr.PerPage,
		PreviousPage: previousPage,
		Page:         pr.Page,
		NextPage:     nextPage,
	}
}
