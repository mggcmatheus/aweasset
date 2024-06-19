package config

import (
	"errors"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
)

type Config struct {
	Debug   string
	Server  ServerConfig
	Cors    CorsConfig
	Datadog DatadogConfig
	Mysql   MysqlConfig
	Mongodb MongodbConfig
}

type DatadogConfig struct {
	Env         string
	ServiceName string
}

type ServerConfig struct {
	Host       string
	Port       string
	ExposePort string
}

type CorsConfig struct {
	AllowOrigins string
}

type MysqlConfig struct {
	ConnectionString string
}

type MongodbConfig struct {
	ConnectionString string
}

func getConfigPath(env string) string {
	if env == "docker" {
		return "config-docker"
	} else if env == "production" {
		return "config-production"
	} else {
		return "config-development"
	}
}

func GetConfig() *Config {
	// Carrega as variáveis de ambiente do arquivo .env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Erro ao carregar arquivo .env: %v", err)
	}

	// Carrega o arquivo de configuração YAML
	cfgPath := getConfigPath(os.Getenv("ENVIRONMENT"))
	v, err := LoadConfig(cfgPath, "yml")
	if err != nil {
		log.Fatalf("Erro ao carregar a configuração %v", err)
	}

	// Substitui as variáveis de ambiente no arquivo de configuração
	replaceEnvVars(v)

	// Parseia a configuração
	cfg, err := ParseConfig(v)
	if err != nil {
		log.Fatalf("Erro ao analisar a configuração %v", err)
	}

	return cfg
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var cfg Config
	err := v.Unmarshal(&cfg)
	if err != nil {
		log.Printf("Não foi possível analisar a configuração: %v", err)
		return nil, err
	}
	return &cfg, nil
}

func LoadConfig(filename string, fileType string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigType(fileType)
	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		log.Printf("Não foi possível ler a configuração: %v", err)
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			return nil, errors.New("arquivo de configuração não encontrado")
		}
		return nil, err
	}
	v.SetDefault("debug", false)
	return v, nil

}

func replaceEnvVars(v *viper.Viper) {
	// Obtem as variáveis de ambiente
	environment := os.Getenv("ENVIRONMENT")
	debug := os.Getenv("DEBUG")

	serverHost := os.Getenv("SERVER_HOST")
	serverPort := os.Getenv("SERVER_PORT")
	serverExposePort := os.Getenv("SERVER_EXPOSE_PORT")

	datadogWithEnv := os.Getenv("DATADOG_WITH_ENV")
	datadogWithService := os.Getenv("DATADOG_WITH_SERVICE")
	datadogWithAddress := os.Getenv("DATADOG_WITH_ADDRESS")

	// Substitui as variáveis no arquivo de configuração
	v.Set("environment", environment)
	v.Set("debug", debug)

	v.Set("server.host", serverHost)
	v.Set("server.port", serverPort)
	v.Set("server.exposePort", serverExposePort)

	v.Set("datadog.env", datadogWithEnv)
	v.Set("datadog.serviceName", datadogWithService)
	v.Set("datadog.address", datadogWithAddress)
}
