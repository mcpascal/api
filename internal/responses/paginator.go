package responses

type Paginator[T interface{}] struct {
	Page  int   `form:"page,default=1" json:"page"`
	Size  int   `form:"size,default=10" json:"size"`
	Data  []T   `json:"data"`
	Total int64 `json:"total"`
}
