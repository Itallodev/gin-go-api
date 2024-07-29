package database

import (
	"github.com/devitallo/gin-go-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaDB() {
	stringDeConexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		panic("Falha ao conectar com o banco de dados: " + err.Error())
	}
	DB.AutoMigrate(&models.Aluno{})
}
