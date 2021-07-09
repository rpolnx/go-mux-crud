package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db * gorm.DB
)

func init() {

	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
}

func Connect() *gorm.DB{	

	dsn := "host=localhost user=" + os.Getenv("DB_USER") + 
	" password=" + os.Getenv("DB_PASS") +  
	" dbname=" + os.Getenv("DB_NAME") + 
	" port=" + os.Getenv("DB_PORT") + " sslmode=disable TimeZone=UTC"
	dbc, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		panic(err)
	}

	db = dbc

	return db
}

func GetDB() *gorm.DB {
	return db
}