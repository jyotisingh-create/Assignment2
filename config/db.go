package config

import (
	"banking-system/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() error {
	_ = godotenv.Load()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_USER", ""),
		getEnv("DB_PASSWORD", ""),
		getEnv("DB_NAME", ""),
		getEnv("DB_PORT", "5432"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	DB = db
	log.Println("Database connected successfully")

	return AutoMigrate()
}

func AutoMigrate() error {
	return DB.AutoMigrate(
		&models.Bank{},
		&models.Branch{},
		&models.Customer{},
		&models.SavingsAccount{},
		&models.CustomerAccount{},
		&models.Transaction{},
		&models.Loan{},
		&models.LoanPayment{},
	)
}

func CloseDB() error {
	db, err := DB.DB()
	if err != nil {
		return err
	}
	return db.Close()
}

func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
