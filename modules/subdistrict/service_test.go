package subdistrict

import (
	"context"
	"testing"
)

type subdistrictRepositoryStub struct {
	findByDistrictCode func(context.Context, string) ([]Subdistrict, error)
	findByQuery        func(context.Context, string, QueryOptions) ([]SubdistrictSearch, error)
	countByQuery       func(context.Context, string) (int, error)
}

func (s subdistrictRepositoryStub) FindByDistrictCode(ctx context.Context, code string) ([]Subdistrict, error) {
	return s.findByDistrictCode(ctx, code)
}

func (s subdistrictRepositoryStub) FindByQuery(ctx context.Context, query string, options QueryOptions) ([]SubdistrictSearch, error) {
	return s.findByQuery(ctx, query, options)
}

func (s subdistrictRepositoryStub) CountByQuery(ctx context.Context, query string) (int, error) {
	return s.countByQuery(ctx, query)
}

func TestServiceFindByQueryCalculatesOffset(t *testing.T) {
	var received QueryOptions
	service := newService(subdistrictRepositoryStub{
		findByDistrictCode: func(context.Context, string) ([]Subdistrict, error) {
			return nil, nil
		},
		findByQuery: func(_ context.Context, _ string, options QueryOptions) ([]SubdistrictSearch, error) {
			received = options
			return []SubdistrictSearch{{Code: "1101010"}}, nil
		},
		countByQuery: func(context.Context, string) (int, error) {
			return 1, nil
		},
	})

	results, count, err := service.FindByQuery(context.Background(), "jakarta", QueryOptions{
		Page:  3,
		Limit: 10,
		Sort:  "asc",
	})
	if err != nil {
		t.Fatalf("FindByQuery returned error: %v", err)
	}
	if received.Offset != 20 {
		t.Fatalf("repository received offset %d, want 20", received.Offset)
	}
	if len(results) != 1 || count != 1 {
		t.Fatalf("unexpected result: %#v, count %d", results, count)
	}
}
