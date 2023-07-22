package config

import (
	"bookapi/helper"
	"fmt"

	// postgres golang driver
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Sarker1724@"
	dbName   = "studentdb"
)

func DatabaseConnection() *gorm.DB {

	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	helper.PanicError(err)

	return db

}
