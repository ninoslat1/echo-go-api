package configs

import (
	// "fmt"

	"echo-api/models"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func RunDatabase(log *logrus.Logger, dbName string) (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Warn("Error loading .env file, using default environment variables")
	}

	cfg := &models.DBConfig{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Port:     os.Getenv("DB_PORT"),
	}

	dsn := fmt.Sprintf(
		// ":cog938gb18@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, dbName,
		// dbName,
	)

	// Open database connection
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Error("Failed to connect to database:", err)
		return nil, err
	}

	log.Infof("Database '%s' connected successfully", dbName)
	return db, nil
}
