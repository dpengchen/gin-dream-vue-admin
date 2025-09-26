package core

type BaseQuery struct {
	Page int `json:"page" form:"page,default=1"`
	Size int `json:"size" form:"size,default=10"`
}
