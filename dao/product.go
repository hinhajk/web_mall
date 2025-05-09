package dao

import (
	"context"
	"gorm.io/gorm"
	"web_mall/models"
	//"web_mall/models"
)

type ProductDao struct {
	*gorm.DB
}

func NewProductDao(ctx context.Context) *ProductDao {
	return &ProductDao{NewDBClient(ctx)}
}

func (dao *ProductDao) CreateProduct(product *models.Product) (err error) {
	return dao.DB.Model(&models.Product{}).Create(&product).Error
}
