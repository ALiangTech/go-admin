package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	HOST     = "139.196.89.94"
	PORT     = 5433
	USER     = "goadmin"
	PASSWORD = "1234@qwer"
	NAME     = "db8eb6e724c4644762bb7c1eb08c023f88admins"
	MAXIDLE  = 10
	MAXOPEN  = 100
)

var DB *gorm.DB
var DSN = fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", HOST, PORT, USER, NAME, PASSWORD)

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
