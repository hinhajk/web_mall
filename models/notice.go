package models

import "gorm.io/gorm"

// Notice 公告模型
type Notice struct {
	gorm.Model
	Text string `gorm:"type:text"`
}
