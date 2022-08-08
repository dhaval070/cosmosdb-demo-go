package main

import "cosmosdb-demo/pkg/di"

func main() {

	app := di.InitializeAPI()
	app.Run()
}
