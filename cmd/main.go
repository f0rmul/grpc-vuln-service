package main

import (
	"log"

	"github.com/f0rmul/vuln-service/config"
	"github.com/f0rmul/vuln-service/internal/server"
	"github.com/f0rmul/vuln-service/pkg/logger"
)

func main() {
	log.Println("Initializing components")

	cfg, err := config.NewConfig("config/config.yml")

	if err != nil {
		log.Fatal(err)
	}

	appLogger := logger.NewApiLogger(cfg)
	appLogger.InitLogger()

	s := server.NewVulnServer(appLogger, cfg)
	appLogger.Fatal(s.Run())
}
