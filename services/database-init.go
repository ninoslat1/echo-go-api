package services

import (
	"database/sql"
	"echo-api/models"
	"fmt"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

// InitDatabase mengecek apakah database sudah ada, jika tidak akan dibuat
func InitDatabase(cfg *models.DBConfig, log *logrus.Logger) {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Koneksi sementara ke PostgreSQL tanpa database
	dsnTemp := fmt.Sprintf(
		"host=127.0.0.1 user=postgres password=cog938gb18 port=5432 sslmode=disable",
		// cfg.Host, cfg.User, cfg.Password, cfg.Port,
	)

	sqlDB, err := sql.Open("postgres", dsnTemp)
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL:", err)
	}
	defer sqlDB.Close()

	var exists bool
	query := fmt.Sprintf("SELECT EXISTS (SELECT FROM pg_database WHERE datname = '%s')", cfg.Name)
	err = sqlDB.QueryRow(query).Scan(&exists)
	if err != nil {
		log.Fatal("Failed to check if database exists:", err)
	}

	if !exists {
		_, err = sqlDB.Exec("CREATE DATABASE %s", cfg.Name)
		if err != nil {
			log.Fatal("Failed to create database. Reason :", err.Error())
		}
		log.Info("Master database created successfully")
	}
}
