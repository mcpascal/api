package models

type User struct {
	Model
	Email    string `gorm:"type:varchar(20);not null;unique" json:"email"`
	Username string `gorm:"type:varchar(20)" json:"username"`
	Nickname string `json:"nickname"`
	Mobile   string `json:"mobile"`
	Password string `json:"-"`
	Status   int    `json:"status"`
	IsAdmin  int    `json:"is_admin"`
	Avatar   string `json:"avatar"`
	Birthday string `json:"birthday"`
	Gender   int    `json:"gender"`
	Address  string `json:"address"`
	Website  string `json:"website"`
	Intro    string `json:"intro"`
}

func NewUser() *User {
	return &User{}
}
