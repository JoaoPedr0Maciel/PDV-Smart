// database/database.go
package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgres() (*gorm.DB, error) {
	dsn := "host=172.18.0.3 port=5432 user=postgres password=postgres dbname=pdv_smart sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error occurred while connecting to database: (%v)", err)
		return nil, err
	}

	return db, nil
}


