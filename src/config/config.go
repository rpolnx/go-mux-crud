package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db * gorm.DB
)

func Connect() *gorm.DB{
	// Please define your user name and password for my sql.

	dsn := "host=localhost user=postgres password=password dbname=postgres port=5432 sslmode=disable TimeZone=UTC"
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