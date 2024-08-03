package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	fmt.Println("Iniciando tentativa de conexão com o banco de dados")

	stringDeConexao := "host=localhost user=postgres password=pass dbname=postgres port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	fmt.Println("Banco de dados conectado com sucesso!")
}
