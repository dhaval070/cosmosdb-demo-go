package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
	"github.com/google/wire"

	"cosmosdb-demo/domain"
	"cosmosdb-demo/pkg/logger"
	"cosmosdb-demo/pkg/repository/interfaces"
)

type repo struct {
	client *azcosmos.Client
	logger.ILogger
}

func NewRepo(c *azcosmos.Client, log logger.ILogger) *repo {
	return &repo{
		c,
		log,
	}
}

func (r *repo) GetFamily(id string) (*domain.Family, error) {
	c, err := r.client.NewContainer("ToDoList", "families")

	if err != nil {
		r.Error(err)
		return nil, fmt.Errorf("%v", err)
	}

	pk := azcosmos.NewPartitionKeyString(id)

	item, err := c.ReadItem(context.Background(), pk, id, nil)

	if err != nil {
		r.Error(err)
		return nil, fmt.Errorf("%v", err)
	}
	log.Println(string(item.Value))

	var family domain.Family
	err = json.Unmarshal(item.Value, &family)

	if err != nil {
		return nil, err
	}

	return &family, nil
}

func (r *repo) Hello() string {
	return "hi from repo"
}

var Wired = wire.NewSet(NewRepo, wire.Bind(new(interfaces.Repo), new(*repo)))
