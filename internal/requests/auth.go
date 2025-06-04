package requests

type Login struct {
	Email      string `form:"username" json:"email" binding:"required,email"`
	Password   string `form:"password" json:"password" binding:"required,min=8,max=56,containsany=0123456789,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ"`
	RememberMe bool   `form:"remember_me" json:"remember_me"`
	Captcha    string `form:"captcha" json:"captcha" binding:"required"`
	SmsCode    string `form:"sms_code" json:"sms_code" binding:"required"`
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
