package main

import (
	"os"

	"github.com/lucasd-coder/aplicacao-estudos-consumer-queue-v2/config"
	"github.com/lucasd-coder/aplicacao-estudos-consumer-queue-v2/internal/app"
	"github.com/lucasd-coder/aplicacao-estudos-consumer-queue-v2/pkg/logger"

	"github.com/ilyakaznacheev/cleanenv"
)

var cfg config.Config

func main() {	
	profile := os.Getenv("GO_PROFILE")
	var path string

	switch profile {
	case "dev":
		path = "./config/config-dev.yml"		
	}

	err := cleanenv.ReadConfig(path, &cfg)
	if err != nil {
		logger.Log.Fatalf("Config error: %v", err)
	}
	config.ExportConfig(&cfg)

	app.Run(&cfg)
}
