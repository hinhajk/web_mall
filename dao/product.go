package dao

import (
	"context"
	"gorm.io/gorm"
	"web_mall/models"
)

type ProductDao struct {
	*gorm.DB
}

func NewProductDao(ctx context.Context) *ProductDao {
	return &ProductDao{NewDBClient(ctx)}
}

func NewProductDaoByDB(db *gorm.DB) *ProductDao {
	return &ProductDao{db}
}

func (dao *ProductDao) CreateProduct(product *models.Product) (err error) {
	return dao.DB.Model(&models.Product{}).Create(&product).Error
}

func (dao *ProductDao) CountProductByCategory(condition map[string]interface{}) (count int64, err error) {
	err = dao.DB.Model(&models.Product{}).Where(condition).Count(&count).Error
	return
}

func (dao *ProductDao) ListCountProductByCategory(condition map[string]interface{}, page models.BasePage) (products []*models.Product, err error) {
	err = dao.DB.Where(condition).Offset((page.PageNum - 1) * page.PageSize).Limit(page.PageSize).Find(&products).Error
	return
}
