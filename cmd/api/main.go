package main

import (
	"cosmosdb-demo/pkg/config"
	"cosmosdb-demo/pkg/di"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
	"github.com/kelseyhightower/envconfig"
)

func initDB(cfg config.Config) *azcosmos.Client {
	cred, err := azcosmos.NewKeyCredential(cfg.DB_KEY)
	if err != nil {
		log.Fatal(err)
	}
	client, err := azcosmos.NewClientWithKey(cfg.DB_ENDPOINT, cred, nil)
	if err != nil {
		log.Fatal(err)
	}

	var n = 1
	n = n

	return client
}

func main() {
	var cfg config.Config

	envconfig.MustProcess("cosmosapp", &cfg)
	log.Println(cfg)

	app := di.InitializeAPI(initDB(cfg))
	app.Run()
}
