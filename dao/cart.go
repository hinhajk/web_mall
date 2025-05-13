package dao

import (
	"context"
	"gorm.io/gorm"
	"web_mall/models"
)

type CartDao struct {
	*gorm.DB
}

func NewCartDao(ctx context.Context) *CartDao {
	return &CartDao{NewDBClient(ctx)}
}

func NewCartDaoByDB(db *gorm.DB) *CartDao {
	return &CartDao{db}
}

func (dao *CartDao) CreateCart(cart *models.Cart) (err error) {
	return dao.DB.Model(&models.Cart{}).Create(cart).Error
}

func (dao *CartDao) ShowCart(uid uint, cId uint) (carts *models.Cart, err error) {
	err = dao.DB.Model(&models.Cart{}).Where("id = ?", cId).Find(&carts).Error
	return
}

func (dao *CartDao) UpdateCart(uid uint, cId uint, cart *models.Cart) (err error) {
	return dao.DB.Model(&models.Cart{}).Where("user_id = ? AND id = ?", uid, cId).Updates(&cart).Error
}

func (dao *CartDao) DeleteCart(uid uint, cId uint) (err error) {
	return dao.DB.Model(&models.Cart{}).Where("user_id = ? AND id = ?", uid, cId).Delete(&models.Cart{}).Error
}

func (dao *CartDao) DeleteAllCarts(uid uint) (err error) {
	return dao.DB.Model(&models.Cart{}).Where("user_id = ?", uid).Delete(&models.Cart{}).Error
}
