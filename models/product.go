package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ProductName   string
	Category      uint
	Title         string
	Info          string
	ImgPath       string
	Price         string
	DiscountPrice string
	OnSale        bool `gorm:"default:true"`
	Num           int  `gorm:"default:0"`
	BossId        uint
	BossName      string
	BossAvatar    string
}
