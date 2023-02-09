package main

import (
	"github.com/bruma1994/go-production/internal/app"
	"github.com/bruma1994/go-production/internal/config"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("config initializing...")
	cfg := config.GetConfig()

	log.Info("logger initializing")

	logger := log.New()

	a, err := app.NewApp(cfg, logger)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Println("running application")
	a.Run()
}
