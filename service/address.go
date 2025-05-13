package service

import (
	"context"
	"strconv"
	"web_mall/dao"
	"web_mall/models"
	"web_mall/pkg/e"
	"web_mall/pkg/utils"
	"web_mall/serializer"
)

type AddressService struct {
	Name    string `json:"name" form:"name"`
	Phone   string `json:"phone" form:"phone"`
	Address string `json:"address" form:"address"`
}

// Create 创建地址
func (service *AddressService) Create(context context.Context, uid uint) serializer.Response {
	addressDao := dao.NewAddressDao(context)
	code := e.Success
	// 地址不考虑去重
	address := &models.Address{
		UserID:  uid,
		Name:    service.Name,
		Phone:   service.Phone,
		Address: service.Address,
	}
	err := addressDao.CreateAddress(address)
	if err != nil {
		code = e.Error
		utils.LogObj.Infoln("err:", err)
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
			Error:   err.Error(),
		}
	}
	return serializer.Response{
		Status:  code,
		Message: e.GetMsg(code),
		Data:    serializer.BuildAddress(address),
	}
}

// Get 查询单个地址
func (service *AddressService) Get(context context.Context, uid uint, id string) serializer.Response {
	addressDao := dao.NewAddressDao(context)
	code := e.Success
	aid, _ := strconv.Atoi(id)
	address, err := addressDao.GetAddressByID(uid, uint(aid))
	if err != nil {
		code = e.Error
		utils.LogObj.Infoln("err:", err)
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
			Error:   err.Error(),
		}
	}
	return serializer.Response{
		Status:  code,
		Message: e.GetMsg(code),
		Data:    serializer.BuildAddress(address),
	}
}

// List 查询用户所有地址
func (service *AddressService) List(context context.Context, uid uint) serializer.Response {
	addressDao := dao.NewAddressDao(context)
	code := e.Success
	address, err := addressDao.GetAddressByUid(uid)
	if err != nil {
		code = e.Error
		utils.LogObj.Infoln("err:", err)
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
			Error:   err.Error(),
		}
	}
	return serializer.Response{
		Status:  code,
		Message: e.GetMsg(code),
		Data:    serializer.BuildAddressS(address),
	}
}

// Update 更新单个地址
func (service *AddressService) Update(context context.Context, uid uint, id string) serializer.Response {
	code := e.Success
	addressDao := dao.NewAddressDao(context)
	aid, _ := strconv.Atoi(id)
	//先找到要更新的记录
	address, err := addressDao.GetAddressByID(uid, uint(aid))
	if err != nil {
		code = e.Error
		utils.LogObj.Infoln("err:", err)
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
			Error:   err.Error(),
		}
	}
	if service.Name != "" {
		address.Name = service.Name
	}
	if service.Phone != "" {
		address.Phone = service.Phone
	}
	if service.Address != "" {
		address.Address = service.Address
	}
	err = addressDao.UpdateAddress(uid, uint(aid), address)
	if err != nil {
		code = e.Error
		utils.LogObj.Infoln("err:", err)
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
			Error:   err.Error(),
		}
	}
	return serializer.Response{
		Status:  code,
		Message: e.GetMsg(code),
		Data:    serializer.BuildAddress(address),
	}
}

// Delete 删除单个地址
func (service *AddressService) Delete(context context.Context, uid uint, id string) serializer.Response {
	code := e.Success
	addressDao := dao.NewAddressDao(context)
	aid, _ := strconv.Atoi(id)
	err := addressDao.DeleteAddress(uint(aid), uid)
	if err != nil {
		code = e.Error
		utils.LogObj.Infoln("err:", err)
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status:  code,
		Message: e.GetMsg(code),
	}
}
