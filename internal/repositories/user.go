package repositories

import (
	"api/internal/models"
	"api/pkg/mysql"
	"fmt"
)

type User struct {
	model *models.User
}

func NewUser() *User {
	return &User{
		model: &models.User{},
	}
}

func (u *User) Create(user *models.User) (*models.User, error) {
	user.Password = "test12312321312"
	if err := mysql.Database.Create(&user).Error; err != nil {
		return &models.User{}, err
	}
	return user, nil
}

func (u *User) DeleteById(id int) error {
	if err := mysql.Database.Where("id = ?", id).Delete(&u.model).Error; err != nil {
		return err
	}
	return nil
}

func (u *User) FindByEmail(email string) (models.User, error) {
	user := models.User{}
	if err := mysql.Database.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (u *User) FindById(id int) (*models.User, error) {
	user := &models.User{}
	if err := mysql.Database.Where("id =?", id).First(user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (u *User) FindByPaginator(page int, limit int) (int64, []models.User, error) {
	var total int64
	users := []models.User{}
	offset := (page - 1) * limit
	if err := mysql.Database.Offset(offset).Limit(limit).Find(&users).Count(&total).Error; err != nil {
		return total, users, err
	}
	if err := mysql.Database.Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		return 0, users, err
	}
	return total, users, nil
}

func (u *User) Search(page int, limit int, conditions []interface{}, orders []interface{}) (int64, []models.User, error) {
	var total int64
	users := []models.User{}
	offset := (page - 1) * limit
	query := mysql.Database.Offset(offset).Limit(limit)
	for _, condition := range conditions {
		fmt.Println(condition)
		// query = query.Where(condition[0], condition[1], condition[2])
	}
	for _, order := range orders {
		fmt.Println(order)
		// query = query.Order(order[0], order[1])
	}
	if err := query.Find(&users).Count(&total).Error; err != nil {
		return total, users, err
	}
	if err := query.Find(&users).Error; err != nil {
		return 0, users, err
	}
	return total, users, nil
}

func (u *User) Update(user *models.User) error {
	if err := mysql.Database.Save(&user).Error; err != nil {
		return err
	}
	return nil
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
