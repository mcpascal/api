package requests

type Paginator struct {
	Page int `form:"page,default=1" binding:"required"`
	Size int `form:"size,default=10" binding:"required"`
}
