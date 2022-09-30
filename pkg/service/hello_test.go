package service

import (
	"cosmosdb-demo/domain"
	"testing"
)

type MockedRepo struct {
}

func (r *MockedRepo) Hello() string {
	return "hello"
}

func (r *MockedRepo) GetFamily(id string) (*domain.Family, error) {
	return nil, nil
}

func TestGetFamily(t *testing.T) {
	r := &MockedRepo{}
	h := hello{r}

	resp := h.Hello()
	if resp != "hello" {
		t.Error("not hello")
	}
}
