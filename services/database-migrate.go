package services

import (
	"echo-api/models"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB, log *logrus.Logger) {
	err := db.AutoMigrate(&models.User{})

	if err != nil {
		log.Fatal("Failed to migrate database", err.Error())
	}
	log.Info("Database migrate successfully")
}
