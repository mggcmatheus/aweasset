package config

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"sync"
	"time"
)

var (
	onceMongoDb sync.Once
	client      *mongo.Client
	errMongoDb  error
)

func GetMongoClient() (*mongo.Client, error) {
	onceMongoDb.Do(func() {
		clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_CONNECTION_STRING"))
		client, errMongoDb = mongo.Connect(context.Background(), clientOptions)
		if errMongoDb != nil {
			fmt.Printf("[ erro ] erro ao conectar ao MongoDB: %v\n", errMongoDb)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		errMongoDb = client.Ping(ctx, nil)
		if errMongoDb != nil {
			err := client.Disconnect(context.Background())
			fmt.Printf("[ erro ] Conexão com o MongoDB: %v\n", err)
			return
		}

		fmt.Println("[ OK ] Conexão com o MongoDB")
	})
	if client == nil {
		return nil, fmt.Errorf("[ erro ] Cliente MongoDB não inicializado")
	}
	return client, nil
}

func CloseMongoClient() {
	if client != nil {
		err := client.Disconnect(context.Background())
		if err != nil {
			fmt.Printf("[ erro ] erro ao desconectar do MongoDB: %v\n", err)
		} else {
			fmt.Println("[ OK ] Conexão com o MongoDB encerrada.")
		}
	}
}
