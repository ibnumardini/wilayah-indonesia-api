package httpx

import (
	"log"
	"net/http"
)

func InternalError(w http.ResponseWriter, operation string, err error) {
	log.Printf("http operation %s failed: %v", operation, err)
	ResponseError(w, Response{
		Status: http.StatusInternalServerError,
	})
}
