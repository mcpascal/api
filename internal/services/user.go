package services

import (
	"api/internal/models"
	"api/internal/repositories"
	"api/internal/requests"
	"api/internal/responses"
)

type User struct {
	repoisitory *repositories.User
}

func NewUser() *User {
	return &User{
		repoisitory: repositories.NewUser(),
	}
}

func (u *User) Index(req *requests.Paginator) (responses.Paginator[responses.User], error) {
	resp := responses.Paginator[responses.User]{}
	data := []responses.User{}
	total, users, err := u.repoisitory.FindByPaginator(req.Page, req.Size)
	if err != nil {
		return resp, err
	}
	resp.Total = total
	for _, user := range users {
		data = append(data, responses.User{
			Id:    int64(user.ID),
			Email: user.Email,
		})
	}
	resp.Data = data
	resp.Page = req.Page
	resp.Size = req.Size
	return resp, nil
}

func (u *User) Show(id int) (responses.IResponse, error) {
	resp := responses.User{}
	user, err := u.repoisitory.FindById(id)
	if err != nil {
		return nil, err
	}
	resp.Id = int64(user.ID)
	resp.Email = user.Email
	return resp, nil
}

func (u *User) Store(req *requests.User) (responses.User, error) {
	resp := responses.User{}
	data := &models.User{
		Email: req.Email,
	}
	user, err := u.repoisitory.Create(data)
	if err != nil {
		return responses.User{}, err
	}
	resp.Id = int64(user.ID)
	resp.Email = user.Email
	return resp, nil
}

func (u *User) Update(id int, req *requests.User) (responses.User, error) {
	resp := responses.User{}
	user, err := u.repoisitory.FindById(id)
	if err != nil {
		return responses.User{}, err
	}

	user.Email = req.Email
	if err := u.repoisitory.Update(user); err != nil {
		return responses.User{}, err
	}
	resp.Id = int64(user.ID)
	resp.Email = user.Email
	return resp, nil
}

func (u *User) Destroy(id int) error {
	_, err := u.repoisitory.FindById(id)
	if err != nil {
		return err
	}
	return u.repoisitory.DeleteById(id)
}
