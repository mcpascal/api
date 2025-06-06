package requests

type Paginator struct {
	Page int `form:"page,default=1" json:"page,default=1" binding:"required"`
	Size int `form:"size,default=10" json:"size,default=10" binding:"required"`
}

type Search struct {
	Paginator
	Orders     []Order    `form:"orders" json:"orders"`
	Conditions Conditions `form:"conditions" json:"conditions"`
}

type Order struct {
	Column string `form:"column" json:"column"`
	Sort   string `form:"sort" json:"sort"`
}

type Conditions struct {
	Relation string      `form:"relation" json:"relation"`
	Items    []Condition `form:"items" json:"items"`
}
type Condition struct {
	Column string      `form:"column" json:"column"`
	Value  interface{} `form:"value" json:"value"`
	Op     string      `form:"op" json:"op"`
}
