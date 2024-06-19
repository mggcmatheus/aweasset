package config

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"
)

var (
	onceMysql sync.Once
	db        *sql.DB
	errMysql  error
)

func GetMySQLConnection(cfg MysqlConfig) (*sql.DB, error) {
	onceMysql.Do(func() {

		// Conecta ao banco de dados MySQL
		db, errMysql = sql.Open("mysql", cfg.ConnectionString)
		if errMysql != nil {
			fmt.Printf("[ erro ] erro ao conectar ao banco de dados MySQL: %v\n", errMysql)
			return
		}

		// Testa a conexão
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		errMysql = db.PingContext(ctx)
		if errMysql != nil {
			db.Close()
			fmt.Printf("[ erro ] erro ao testar a conexão com o MySQL: %v\n", errMysql)
			return
		}

		fmt.Println("[ OK ] Conexão com o MySQL estabelecida com sucesso.")
	})
	if db == nil {
		return nil, fmt.Errorf("[ erro ] conexão com o MySQL não inicializada")
	}
	return db, nil
}

func CloseMySQLConnection() {
	if db != nil {
		err := db.Close()
		if err != nil {
			fmt.Printf("[ erro ] erro ao fechar a conexão com o MySQL: %v\n", err)
		} else {
			fmt.Println("[ OK ] Conexão com o MySQL encerrada.")
		}
	}
}
