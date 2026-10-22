package province

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi"
)

type provinceServiceStub struct {
	findAll    func(context.Context) ([]Province, error)
	findByCode func(context.Context, string) (*Province, error)
}

func (s provinceServiceStub) FindAll(ctx context.Context) ([]Province, error) {
	return s.findAll(ctx)
}

func (s provinceServiceStub) FindByCode(ctx context.Context, code string) (*Province, error) {
	return s.findByCode(ctx, code)
}

func TestHandlerFindByCodeSuccess(t *testing.T) {
	handler := newHandler(provinceServiceStub{
		findAll: func(context.Context) ([]Province, error) {
			return nil, nil
		},
		findByCode: func(_ context.Context, code string) (*Province, error) {
			return &Province{Code: code, Name: "Aceh"}, nil
		},
	})

	request := httptest.NewRequest(http.MethodGet, "/provinces/11", nil)
	routeContext := chi.NewRouteContext()
	routeContext.URLParams.Add("code", "11")
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, routeContext))
	recorder := httptest.NewRecorder()

	handler.FindByCode(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("FindByCode returned status %d, want %d", recorder.Code, http.StatusOK)
	}

	var response struct {
		Ok     bool     `json:"ok"`
		Result Response `json:"result"`
	}
	if err := json.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if !response.Ok || response.Result.Code != "11" {
		t.Fatalf("unexpected response: %#v", response)
	}
}

func TestHandlerFindByCodeNotFound(t *testing.T) {
	handler := newHandler(provinceServiceStub{
		findAll: func(context.Context) ([]Province, error) {
			return nil, nil
		},
		findByCode: func(context.Context, string) (*Province, error) {
			return nil, nil
		},
	})

	request := httptest.NewRequest(http.MethodGet, "/provinces/99", nil)
	routeContext := chi.NewRouteContext()
	routeContext.URLParams.Add("code", "99")
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, routeContext))
	recorder := httptest.NewRecorder()

	handler.FindByCode(recorder, request)

	if recorder.Code != http.StatusNotFound {
		t.Fatalf("FindByCode returned status %d, want %d", recorder.Code, http.StatusNotFound)
	}
}

func TestHandlerFindByCodeInternalError(t *testing.T) {
	handler := newHandler(provinceServiceStub{
		findAll: func(context.Context) ([]Province, error) {
			return nil, nil
		},
		findByCode: func(context.Context, string) (*Province, error) {
			return nil, errors.New("database failure")
		},
	})

	request := httptest.NewRequest(http.MethodGet, "/provinces/11", nil)
	routeContext := chi.NewRouteContext()
	routeContext.URLParams.Add("code", "11")
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, routeContext))
	recorder := httptest.NewRecorder()

	handler.FindByCode(recorder, request)

	if recorder.Code != http.StatusInternalServerError {
		t.Fatalf("FindByCode returned status %d, want %d", recorder.Code, http.StatusInternalServerError)
	}
	if recorder.Body.Len() == 0 || strings.Contains(recorder.Body.String(), "database failure") {
		t.Fatal("internal error details should not be present in response")
	}
}
