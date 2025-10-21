package db

import (
	"fmt"
	calculationservice "main/internal/calculationService"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=yourpassword dbname=postgres port=5432 sslmode=disable"
	var err error

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("could not connect to database: %v", err)
	}

	if err := db.AutoMigrate(&calculationservice.Calculation{}); err != nil {
		return nil, fmt.Errorf("could not migrates: %v", err)
	}
	return db, nil
}
