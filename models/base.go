package models

// BasePage Base 进行分页操作
type BasePage struct {
	PageNum  int `form:"pageNum" json:"pageNum"`
	PageSize int `form:"pageSize" json:"pageSize"`
}
