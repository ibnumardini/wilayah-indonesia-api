package pagination

type Metadata struct {
	Total        int `json:"total"`
	PerPage      int `json:"per_page"`
	PreviousPage int `json:"previous_page"`
	Page         int `json:"page"`
	NextPage     int `json:"next_page"`
}

func (m Metadata) ToMap() map[string]interface{} {
	previousPage := 0
	if m.Page > 1 {
		previousPage = m.Page - 1
	}

	nextPage := 0
	if m.Page*m.PerPage < m.Total {
		nextPage = m.Page + 1
	}

	return map[string]interface{}{
		"total":         m.Total,
		"per_page":      m.PerPage,
		"previous_page": previousPage,
		"page":          m.Page,
		"next_page":     nextPage,
	}
}
