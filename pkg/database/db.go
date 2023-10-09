package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB


func GetDB() *gorm.DB {
	return db
}

func InitDB() {
	loadErr := godotenv.Load()
	if loadErr != nil {
		log.Fatal("error leading file .env")
	}

	DBHOST := os.Getenv("DB_HOST")
	DBPORT := os.Getenv("DB_PORT")
	DBUSER := os.Getenv("DB_USER")
	DBPASS := os.Getenv("DB_PASSWORD")
	DBNAME := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", DBHOST,DBUSER, DBPASS, DBNAME, DBPORT)
	var gormErr error 
	db, gormErr = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if gormErr != nil {
		panic((gormErr.Error()))
	}

	sqlDB, sqlDBErr := db.DB()
	if sqlDBErr != nil {
		panic(sqlDBErr.Error())
	}

	pingErr := sqlDB.Ping()
	if pingErr != nil {
		panic(pingErr.Error())
	}

	fmt.Println("success connect to database")
	
}

