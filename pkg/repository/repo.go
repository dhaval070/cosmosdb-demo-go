package repository

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
	"github.com/google/wire"

	"cosmosdb-demo/domain"
	"cosmosdb-demo/pkg/repository/interfaces"
)

type repo struct {
	client *azcosmos.Client
}

func NewRepo(c *azcosmos.Client) *repo {
	return &repo{
		client: c,
	}
}

func (r *repo) GetFamily(id string) (*domain.Family, error) {
	c, err := r.client.NewContainer("dhaval", "families")

	if err != nil {
		return nil, err
	}

	pk := azcosmos.NewPartitionKeyString(id)

	item, err := c.ReadItem(context.Background(), pk, id, nil)

	if err != nil {
		return nil, err
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
