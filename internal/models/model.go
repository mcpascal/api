package models

import (
	"os"

	"gorm.io/gorm"
)

type Model struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt int64          `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt int64          `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

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
