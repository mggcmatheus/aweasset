package main

import (
	"aweasset/internal/infra/api"
	"aweasset/internal/infra/config"
	"fmt"
	"github.com/joho/godotenv"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"log"
)

const banner = `
Version: %s
---------------------
`

var Version = "v1.0.0"

func main() {
	fmt.Printf(banner, Version)
	cfg := config.GetConfig()

	if cfg.Debug == "false" {
		tracer.Start(
			tracer.WithEnv(cfg.Datadog.Env),
			tracer.WithService(cfg.Datadog.ServiceName),
			tracer.WithAgentAddr("datadog-agent:8126"),
		)
		// When the tracer is stopped, it will flush everything it has to the Datadog Agent before quitting.
		defer tracer.Stop()
	}
	// Carrega as variáveis de ambiente do arquivo .env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Erro ao carregar arquivo .env: %v", err)
	}

	api.InitServer(cfg)

}
