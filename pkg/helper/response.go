package helper

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  int                    `json:"-"`
	Ok      bool                   `json:"ok"`
	Message string                 `json:"message"`
	Meta    map[string]interface{} `json:"meta"`
	Result  interface{}            `json:"result"`
}

func (r *Response) SetWatermark() {
	if r.Meta == nil {
		r.Meta = make(map[string]interface{})
	}

	r.Meta["watermark"] = "https://ibnu.mardini.dev"
}

func ResponseSuccess(w http.ResponseWriter, response Response) {
	if response.Status == 0 {
		response.Status = http.StatusOK
	}

	if response.Message == "" {
		response.Message = http.StatusText(response.Status)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.Status)

	var r = Response{
		Ok:      true,
		Message: response.Message,
		Meta:    response.Meta,
		Result:  response.Result,
	}

	r.SetWatermark()

	json.NewEncoder(w).Encode(r)
}

func ResponseError(w http.ResponseWriter, response Response) {
	if response.Status == 0 {
		response.Status = http.StatusInternalServerError
	}

	if response.Message == "" {
		response.Message = http.StatusText(response.Status)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.Status)

	var r = Response{
		Ok:      false,
		Message: response.Message,
		Result:  struct{}{},
	}

	r.SetWatermark()

	json.NewEncoder(w).Encode(r)
}
