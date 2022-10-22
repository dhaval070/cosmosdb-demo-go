package main

import (
	"cosmosdb-demo/pkg/config"
	"cosmosdb-demo/pkg/di"
	"cosmosdb-demo/pkg/logger"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
	"github.com/joho/godotenv"
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

	return client
}

func main() {
	var cfg config.Config

	if err := godotenv.Load("app.env"); err != nil {
		log.Println(err)
	}

	envconfig.MustProcess("cosmosapp", &cfg)

	log := logger.Must(logger.NewLogger())
	log.Debug(cfg)

	app := di.InitializeAPI(initDB(cfg), log)
	app.Run()
}
