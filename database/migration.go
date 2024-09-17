package database

import (
	"sales-api/models"

	"gorm.io/gorm"
)

func AutoMigration(connection *gorm.DB) {
	connection.Debug().AutoMigrate(
		&models.User{},
		&models.Post{},
		&models.Comment{},
	)
}
