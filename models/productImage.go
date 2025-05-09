package models

import "gorm.io/gorm"

type ProductImage struct {
	gorm.Model
	ProductId uint `gorm:"not null"`
	ImgPath   string
}
