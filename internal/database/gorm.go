package database

import (
	"fmt"
	"log"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ServiceGorm interface {
	DB() *gorm.DB
}

type serviceGorm struct {
	db *gorm.DB
}

func (s *serviceGorm) DB() *gorm.DB {
	return s.db
}

var (
	instance *serviceGorm
	once     sync.Once
)

// ConnectDatabase menggunakan pola Singleton agar koneksi database tidak dibuat berulang kali
func ConnectDatabase() (*serviceGorm, error) {
	var err error
	once.Do(func() {
		// Pastikan semua environment variable ada
		host := os.Getenv("DB_HOST")
		user := os.Getenv("DB_USERNAME")
		pass := os.Getenv("DB_PASSWORD")
		dbname := os.Getenv("DB_DATABASE")
		port := os.Getenv("DB_PORT")

		if host == "" || user == "" || pass == "" || dbname == "" || port == "" {
			log.Println("Database configuration is missing required environment variables")
			err = fmt.Errorf("missing database configuration")
			return
		}

		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			host, user, pass, dbname, port,
		)

		DB, dbErr := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if dbErr != nil {
			log.Println("Failed to connect to database:", dbErr)
			err = dbErr
			return
		}

		autoMigrateModels(DB)

		instance = &serviceGorm{db: DB}
	})

	return instance, err
}
