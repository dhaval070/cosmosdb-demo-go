package service

import (
	"cosmosdb-demo/domain"
	"cosmosdb-demo/pkg/repository"
	repo "cosmosdb-demo/pkg/repository/interfaces"
	"cosmosdb-demo/pkg/service/interfaces"

	"github.com/google/wire"
)

type hello struct {
	repo repo.Repo
}

func NewHelloService(repo repo.Repo) *hello {
	return &hello{repo}
}

func (h *hello) Hello() string {
	return h.repo.Hello()
}

func (h *hello) GetFamily(id string) (*domain.Family, error) {
	family, err := h.repo.GetFamily(id)
	if err != nil {
		return nil, err
	}

	return family, nil
}

var Wired = wire.NewSet(repository.Wired, NewHelloService, wire.Bind(new(interfaces.Hello), new(*hello)))
