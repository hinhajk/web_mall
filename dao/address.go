package dao

import (
	"context"
	"gorm.io/gorm"
	"web_mall/models"
)

type AddressDao struct {
	*gorm.DB
}

func NewAddressDao(ctx context.Context) *AddressDao {
	return &AddressDao{NewDBClient(ctx)}
}

func NewAddressDaoByDB(db *gorm.DB) *AddressDao {
	return &AddressDao{db}
}

// CreateAddress 创建操作
func (dao *AddressDao) CreateAddress(address *models.Address) (err error) {
	return dao.DB.Model(&models.Address{}).Create(address).Error
}

// GetAddressByID 查询单个地址
func (dao *AddressDao) GetAddressByID(uid uint, id uint) (address *models.Address, err error) {
	err = dao.DB.Model(&models.Address{}).Where("id = ? AND user_id = ?", id, uid).First(&address).Error
	return
}

// GetAddressByUid 查询用户所有地址
func (dao *AddressDao) GetAddressByUid(uid uint) (address []*models.Address, err error) {
	err = dao.DB.Model(&models.Address{}).Where("user_id = ?", uid).Find(&address).Error
	return
}

// UpdateAddress 更新单个地址
func (dao *AddressDao) UpdateAddress(uid uint, aid uint, address *models.Address) (err error) {
	err = dao.DB.Model(&models.Address{}).Where("id = ? AND user_id = ?", aid, uid).Updates(&address).Error
	return
}

// DeleteAddress 删除单个地址
func (dao *AddressDao) DeleteAddress(aid uint, uid uint) (err error) {
	return dao.DB.Model(&models.Address{}).Where("user_id = ? AND id = ?", uid, aid).Delete(&models.Address{}).Error
}
