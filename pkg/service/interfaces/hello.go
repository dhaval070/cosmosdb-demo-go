package interfaces

import "cosmosdb-demo/domain"

type Hello interface {
	Hello() string
	GetFamily(id string) (*domain.Family, error)
}
