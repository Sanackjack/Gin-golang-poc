package config

import (
	"edge/helper"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "docker"
	dbName   = "postgres"
)

func DatabaseConnection() *gorm.DB {
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable search_path=test ", host, port, user, password, dbName)
	//	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})

	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // Disable table name pluralization
		},
	})

	helper.ErrorPanic(err)

	return db
}
