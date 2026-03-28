package database

import (
	"database/sql"
	"fmt"
	"go-task-manager/internal/config"
	"log"
	"time"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func ConnectDB(cfg *config.Config) (*sql.DB, error) {
	var db *sql.DB
	var err error

	// Retry mekanizması (Docker'da DB geç kalkabilir)
	maxRetries := 5
	for i := 0; i < maxRetries; i++ {
		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.SSLMode)

		db, err = sql.Open("postgres", dsn)
		if err != nil {
			log.Printf("Failed to open valid driver connection: %v", err)
			time.Sleep(2 * time.Second)
			continue
		}

		if err = db.Ping(); err == nil {
			log.Println("Successfully connected to the database!")
			return db, nil
		}

		log.Printf("Failed to ping database (Attempt %d/%d): %v", i+1, maxRetries, err)
		time.Sleep(2 * time.Second)
	}

	return nil, fmt.Errorf("failed to connect to database after %d attempts: %w", maxRetries, err)
}
