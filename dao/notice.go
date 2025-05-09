package dao

import (
	"context"
	"gorm.io/gorm"
	"web_mall/models"
)

type NoticeDao struct {
	*gorm.DB
}

func NewNoticeDao(ctx context.Context) *UserDao {
	return &UserDao{NewDBClient(ctx)}
}

// FindNoticeById 根据id获取notice
func (dao *UserDao) FindNoticeById(id uint) (*models.Notice, error) {
	notice := &models.Notice{}
	err := dao.DB.Model(&models.Notice{}).Where("id = ?", id).First(&notice).Error
	return notice, err
}
