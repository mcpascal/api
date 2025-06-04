package repositories

import "api/internal/models"

type Auth struct {
	user *models.User
}

func NewAuth() *Auth {
	return &Auth{
		user: models.NewUser(),
	}
}

func (a *Auth) Login() {

}
