package models

import (
	"os"

	"gorm.io/gorm"
)

func InitializeTables(database *gorm.DB) {
	err := database.Set("gorm:table_options", "CHARSET=utf8mb4").AutoMigrate(
	// &App{},
	// &User{},
	// &Journal{},
	)
	if err != nil {
		os.Exit(0)
	}
}
