package repositories

import (
	"api/internal/models"
	"api/pkg/mysql"
)

type User struct {
	model models.User
}

func NewUser() *User {
	return &User{
		model: models.User{},
	}
}

func (u *User) FindByEmail(email string) (models.User, error) {
	user := models.User{}
	if err := mysql.Database.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

// func (u *User) Create(user *models.User) (models.User, error) {
// 	if err := mysql.Database.Create(&user).Error; err != nil {
// 		return models.User{}, err
// 	}
// 	return user, nil
// }

// func (u *User) HashPassword(password string) (string, error) {
// 	return password, nil
// }
