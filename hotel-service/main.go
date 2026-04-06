package main

import (
	"hotel-service/app"
	env "hotel-service/config/env"
)

func main() {
	println("Running GO App")
	env.Load()
	appConfig := app.NewConfig()
	appServer := app.NewApplication(appConfig)

	app.Run(appServer)
}
