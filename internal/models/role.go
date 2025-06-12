package models

type Role struct {
	Model
	Name string `gorm:"type:varchar(20);not null;unique" json:"name"`
}
