package dao

import (
	"context"
	"gorm.io/gorm"
	"web_mall/models"
	//"web_mall/models"
)

type ProductImageDao struct {
	*gorm.DB
}

func NewProductImageDao(ctx context.Context) *ProductImageDao {
	return &ProductImageDao{NewDBClient(ctx)}
}

func NewProductImageDaoByDB(db *gorm.DB) *ProductImageDao {
	return &ProductImageDao{db}
}

func (dao *ProductImageDao) CreateProductImg(productImg *models.ProductImage) (err error) {
	return dao.DB.Model(&models.ProductImage{}).Create(&productImg).Error
}

func (dao *ProductImageDao) ListProductImg(id uint) (productImg []*models.ProductImage, err error) {
	err = dao.DB.Model(&models.ProductImage{}).Where("product_id = ?", id).Find(&productImg).Error
	return
}
