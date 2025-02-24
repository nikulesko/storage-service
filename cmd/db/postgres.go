package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type CurrencyData struct {
	ID       int64   `gorm:"primary_key"`
	Datetime int64   `gorm:"not null"`
	Rate     float32 `gorm:"not null"`
}

type Repository struct {
	db *gorm.DB
}

func NewRepository() (*Repository, error) {
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		return nil, fmt.Errorf("DB_HOST environment variable not set")
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		return nil, fmt.Errorf("DB_PORT environment variable not set")
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		return nil, fmt.Errorf("DB_USER environment variable not set")
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		return nil, fmt.Errorf("DB_PASSWORD environment variable not set")
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		return nil, fmt.Errorf("DB_NAME environment variable not set")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		dbHost, dbUser, dbPassword, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	log.Println("Connected to Postgres database!")

	var tableExists bool
	sqlDB, _ := db.DB()
	existsQuery := fmt.Sprintf("SELECT EXISTS (SELECT FROM information_schema.tables WHERE table_name = '%s')", dbName)
	err = sqlDB.QueryRow(existsQuery).Scan(&tableExists)
	if err != nil {
		return nil, fmt.Errorf("failed to check if table exists: %w", err)
	}

	if !tableExists {
		err = db.AutoMigrate(&CurrencyData{})
		if err != nil {
			return nil, fmt.Errorf("failed to migrate database: %w", err)
		}
		log.Println("Database migration completed!")
	} else {
		log.Printf("Table '%s' already exists, skipping migration\n", dbName)
	}

	return &Repository{db: db}, nil
}

func (r *Repository) SaveCurrencyData(data *CurrencyData) error {
	result := r.db.Create(data)
	if result.Error != nil {
		return fmt.Errorf("failed to create currency data: %w", result.Error)
	}

	return nil
}

func (r *Repository) GetAllRecords() (*[]CurrencyData, error) {
	var data []CurrencyData
	result := r.db.Find(&data)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to get data: %w", result.Error)
	}
	return &data, nil
}
