package interfaces

import "cosmosdb-demo/domain"

type Repo interface {
	Hello() string
	GetFamily(id string) (*domain.Family, error)
}
