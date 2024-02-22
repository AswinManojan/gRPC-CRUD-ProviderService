package db

import (
	"fmt"
	"log"

	"github.com/grpc/gobus/service/config"
	model "github.com/grpc/gobus/service/pkg/Model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect_DB(cfg *config.Config) *gorm.DB {
	host := cfg.HOST
	username := cfg.USERNAME
	password := cfg.PASSWORD
	port := cfg.PORT
	sslmode := cfg.SSLMODE
	dbname := cfg.DBNAME
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, username, password, dbname, port, sslmode)
	Db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to DB")
	}
	Db.AutoMigrate(&model.Provider{})
	return Db
}
