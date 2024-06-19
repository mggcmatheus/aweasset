package usecases

import (
	"boirderplate-go-clean/infra/config"
	"database/sql"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type ExampleUseCase struct {
	mongoClient  *mongo.Client
	clientOracle *sql.DB
}

func NewExampleUseCase() (*ExampleUseCase, error) {
	// Instanciar a conexão do MongoDB
	mongoClient, err := config.GetMongoClient()
	if err != nil {
		return nil, fmt.Errorf("erro ao obter cliente MongoDB: %v", err)
	}

	// Instanciar a conexão do Oracle
	clientOracle, err := config.GetOracleConnection()
	if err != nil {
		return nil, fmt.Errorf("erro ao obter conexão Oracle: %v", err)
	}

	return &ExampleUseCase{
		mongoClient:  mongoClient,
		clientOracle: clientOracle,
	}, nil
}

func (uc *ExampleUseCase) Execute() {
	// Iniciar o tempo de execução
	startTime := time.Now()

	endTime := time.Now()                 // Registra a hora de término
	elapsedTime := endTime.Sub(startTime) // Calcula o tempo decorrido

	fmt.Printf("Sincronização concluída em %s\n", elapsedTime)
	fmt.Printf("Tempo total de sincronização: %s\n", time.Since(startTime))
}

func (uc *ExampleUseCase) CloseConnections() {
	config.CloseMongoClient()
	config.CloseOracleConnection()
}
