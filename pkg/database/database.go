package database

import (
	"asyncfiber/pkg/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() (*gorm.DB, error) {
	dsn, _ := utils.ConnectionURLBuilder("postgres")
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
