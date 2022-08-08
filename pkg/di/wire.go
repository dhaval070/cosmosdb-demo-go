//go:build wireinject

package di

import (
	"cosmosdb-demo/pkg/api"
	"cosmosdb-demo/pkg/handler"

	"github.com/google/wire"
)

func InitializeAPI() *api.Server {
	wire.Build(handler.Wired, api.NewServer)
	return &api.Server{}
}
