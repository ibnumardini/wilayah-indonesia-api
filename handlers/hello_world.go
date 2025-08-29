package handlers

import (
	"net/http"

	"github.com/go-chi/render"
)

type HelloWorldHandler struct{}

func NewHelloWorldHandler() *HelloWorldHandler {
	return &HelloWorldHandler{}
}

func (h *HelloWorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, map[string]string{"hello": "world"})
}
