package database

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() error {
	_, err := InitPostgres()
	if err != nil {
		fmt.Printf("Error trying initialize postgres db")
		return err
	}

	return nil
}


func GetDB() *gorm.DB {
	if db == nil {
		log.Fatal("Database connection is not initialized")
	}
	return db
}