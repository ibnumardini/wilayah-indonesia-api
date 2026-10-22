package province

import (
	"context"
	"errors"
	"testing"
)

type provinceRepositoryStub struct {
	findAll    func(context.Context) ([]Province, error)
	findByCode func(context.Context, string) (*Province, error)
}

func (s provinceRepositoryStub) FindAll(ctx context.Context) ([]Province, error) {
	return s.findAll(ctx)
}

func (s provinceRepositoryStub) FindByCode(ctx context.Context, code string) (*Province, error) {
	return s.findByCode(ctx, code)
}

func TestServiceFindAll(t *testing.T) {
	want := []Province{{Code: "11", Name: "Aceh"}}
	service := newService(provinceRepositoryStub{
		findAll: func(context.Context) ([]Province, error) {
			return want, nil
		},
		findByCode: func(context.Context, string) (*Province, error) {
			return nil, nil
		},
	})

	got, err := service.FindAll(context.Background())
	if err != nil {
		t.Fatalf("FindAll returned error: %v", err)
	}
	if len(got) != 1 || got[0] != want[0] {
		t.Fatalf("FindAll returned %#v, want %#v", got, want)
	}
}

func TestServiceFindByCodePropagatesError(t *testing.T) {
	wantErr := errors.New("repository failure")
	service := newService(provinceRepositoryStub{
		findAll: func(context.Context) ([]Province, error) {
			return nil, nil
		},
		findByCode: func(context.Context, string) (*Province, error) {
			return nil, wantErr
		},
	})

	_, err := service.FindByCode(context.Background(), "11")
	if !errors.Is(err, wantErr) {
		t.Fatalf("FindByCode returned error %v, want %v", err, wantErr)
	}
}
