package service

import (
	"cosmosdb-demo/domain"
	"cosmosdb-demo/pkg/logger"
	"cosmosdb-demo/pkg/repository"
	repo "cosmosdb-demo/pkg/repository/interfaces"
	"cosmosdb-demo/pkg/service/interfaces"

	"github.com/google/wire"
)

type hello struct {
	repo repo.Repo
	logger.ILogger
}

func NewHelloService(repo repo.Repo, log logger.ILogger) *hello {
	return &hello{repo, log}
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
