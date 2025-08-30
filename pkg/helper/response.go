package helper

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  int         `json:"-"`
	Ok      bool        `json:"ok"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
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
	json.NewEncoder(w).Encode(Response{
		Ok:      true,
		Message: response.Message,
		Result:  response.Result,
	})
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
	json.NewEncoder(w).Encode(Response{
		Ok:      false,
		Message: response.Message,
		Result:  struct{}{},
	})
}
