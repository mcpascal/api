package requests

type Login struct {
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required,min=8,containsany=0123456789,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ"`
	RememberMe bool   `json:"remember_me"`
}

type Register struct {
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required,min=8,containsany=0123456789,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

type ResetPassword struct {
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required,min=8,containsany=0123456789,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

type RefreshToken struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}
