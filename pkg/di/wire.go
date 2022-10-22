//go:build wireinject

package di

import (
	"cosmosdb-demo/pkg/api"
	"cosmosdb-demo/pkg/handler"

	"cosmosdb-demo/pkg/logger"
	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
	"github.com/google/wire"
)

func InitializeAPI(client *azcosmos.Client, log *logger.Logger) *api.Server {
	wire.Build(handler.Wired, api.NewServer)
	return &api.Server{}
}
