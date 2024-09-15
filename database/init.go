package database

import (
	"fmt"

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
	return db
}