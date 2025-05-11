package models

// BasePage Base 进行分页操作
type BasePage struct {
	PageNum  int `form:"page_num" json:"page_num"`
	PageSize int `form:"page_size" json:"page_size"`
}
