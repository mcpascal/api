package requests

type User struct {
	Email string `form:"email" json:"email" binding:"required,email"`
	// Password string `form:"password" json:"password" binding:"required,min=8,max=56,containsany=0123456789,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ"`
}
