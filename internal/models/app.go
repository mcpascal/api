package models

type App struct {
	Model
	Name        string `gorm:"type:varchar(20);not null;unique" json:"email"`
	Description string `gorm:"type:text" json:"description"`
}
