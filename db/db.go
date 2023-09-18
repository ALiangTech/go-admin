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
	NAME     = "postgres"
	MAXIDLE  = 10
	MAXOPEN  = 100
)

var DB *gorm.DB

func init() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", HOST, PORT, USER, NAME, PASSWORD)
	DBs, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB = DBs
	if err != nil {
		fmt.Print("db start fail")
		panic(err)
	}
	sqlDb, err := DB.DB()
	if err != nil {
		panic(err)
	}
	sqlDb.SetMaxIdleConns(MAXIDLE)
	sqlDb.SetMaxOpenConns(MAXOPEN)
	fmt.Print("db start success")
}
