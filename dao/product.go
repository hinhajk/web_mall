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

func (dao *ProductDao) SearchProduct(info string, basePage models.BasePage) (products []*models.Product, count int64, err error) {
	err = dao.DB.Model(&models.Product{}).
		Where("title LIKE ? OR info LIKE ?", "%"+info+"%", "%"+info+"%").Count(&count).Error
	if err != nil {
		return
	}

	err = dao.DB.Model(&models.Product{}).
		Where("title LIKE ? OR info LIKE ?", "%"+info+"%", "%"+info+"%").
		Offset((basePage.PageNum - 1) * basePage.PageSize).Limit(basePage.PageSize).
		Find(&products).Error
	return
}

func (dao *ProductDao) ShowProductByID(id uint) (product *models.Product, err error) {
	err = dao.DB.Model(&models.Product{}).Where("id = ?", id).First(&product).Error
	return
}

func (dao *ProductDao) DeleteProduct(uid uint, pid uint) (err error) {
	err = dao.DB.Model(&models.Product{}).Where("id = ? AND boss_id = ?", pid, uid).Delete(&models.Product{}).Error
	return
}

func (dao *ProductDao) UpdateById(uid uint, pid uint, product *models.Product) (err error) {
	err = dao.DB.Model(&models.Product{}).Where("id = ? AND boss_id = ?", pid, uid).Updates(&product).Error
	return
}
