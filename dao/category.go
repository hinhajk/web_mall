package dao

import (
	"context"
	"gorm.io/gorm"
	"web_mall/models"
)

type CategoryDao struct {
	*gorm.DB
}

func NewCategoryDao(ctx context.Context) *CategoryDao {
	return &CategoryDao{NewDBClient(ctx)}
}

func NewCategoryDaoByDB(db *gorm.DB) *CategoryDao {
	return &CategoryDao{db}
}

// FindCategoryById 查找分类
func (dao *CategoryDao) FindCategoryById() (category []models.Favorites, err error) {
	err = dao.DB.Model(&models.Favorites{}).Find(&category).Error
	return category, err
}
