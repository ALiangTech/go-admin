package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	HOST     = "localhost"
	PORT     = 5432
	USER     = "core"
	PASSWORD = "hj1234"
	NAME     = "admin"
	MAXIDLE  = 10
	MAXOPEN  = 100
)

var DB *gorm.DB

func init() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", HOST, PORT, USER, NAME, PASSWORD)
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDb, err := DB.DB()
	if err != nil {
		panic(err)
	}
	sqlDb.SetMaxIdleConns(MAXIDLE)
	sqlDb.SetMaxOpenConns(MAXOPEN)

}
