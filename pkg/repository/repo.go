package repository

import (
	"cosmosdb-demo/pkg/repository/interfaces"

	"github.com/google/wire"
)

type repo struct {
}

func NewRepo() *repo {
	return &repo{}
}

func (r *repo) Hello() string {
	return "hi from repo"
}

var Wired = wire.NewSet(NewRepo, wire.Bind(new(interfaces.Repo), new(*repo)))
