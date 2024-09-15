package database

import (
	"log"
	"pdv/schemas"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgres() (*gorm.DB, error){
	dsn := "host=my_postgres_container port=5432 user=postgres password=postgres dbname=pdv_smart sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error occurred to connect to database: (%v)", err)
	}

	db.AutoMigrate(&schemas.User{}, &schemas.Company{})

	return db, nil
	
}
