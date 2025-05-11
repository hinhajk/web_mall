package models

import "gorm.io/gorm"

// 商品分类
type Favorites struct {
	gorm.Model
	CategoryName string
}
