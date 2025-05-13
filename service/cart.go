package service

import (
	"context"
	"fmt"
	"strconv"
	"web_mall/dao"
	"web_mall/models"
	"web_mall/pkg/e"
	"web_mall/pkg/utils"
	"web_mall/serializer"
)

type CartService struct {
	ID        uint `json:"id" form:"id"`
	ProductId uint `json:"product_id" form:"product_id"`
	BossId    uint `json:"boss_id" form:"boss_id"`
	Num       uint `json:"num" form:"num"`
}

// Create 增加购物车
func (service *CartService) Create(ctx context.Context, uid uint) serializer.Response {
	code := e.Success
	//首先判断是否有这个商品
	productDao := dao.NewProductDao(ctx)
	product, err := productDao.ShowProductByID(service.ProductId)
	if err != nil {
		code = e.Error
		utils.LogObj.Infoln("err", err)
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
			Error:   err.Error(),
		}
	}
	//判断老板是否存在
	bossDao := dao.NewUserDao(ctx)
	boss, err := bossDao.FindUserById(service.BossId)
	if err != nil {
		code = e.Error
		utils.LogObj.Infoln("err", err)
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
			Error:   err.Error(),
		}
	}

	cartDao := dao.NewCartDao(ctx)
	cart := &models.Cart{
		UserId:    uid,
		ProductId: product.ID,
		BossId:    boss.ID,
		Num:       service.Num,
		MaxNum:    uint(product.Num),
	}
	err = cartDao.CreateCart(cart)
	if err != nil {
		code = e.Error
		utils.LogObj.Infoln("err", err)
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
			Error:   err.Error(),
		}
	}
	return serializer.Response{
		Status:  code,
		Message: e.GetMsg(code),
		Data:    serializer.BuildCart(cart, product, boss),
	}
}

// Show 查询购物车
func (service *CartService) Show(ctx context.Context, uid uint, id string) serializer.Response {
	code := e.Success
	cartDao := dao.NewCartDao(ctx)
	cId, _ := strconv.Atoi(id)
	carts, err := cartDao.ShowCart(uid, uint(cId))
	if err != nil {
		code = e.Error
		fmt.Println(11111)
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
			Error:   err.Error(),
		}
	}

	//首先判断是否有这个商品
	productDao := dao.NewProductDao(ctx)
	product, err := productDao.ShowProductByID(carts.ProductId)
	if err != nil {
		code = e.Error
		utils.LogObj.Infoln("err", err)
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
			Error:   err.Error(),
		}
	}
	//判断老板是否存在
	bossDao := dao.NewUserDao(ctx)
	boss, err := bossDao.FindUserById(carts.BossId)
	if err != nil {
		code = e.Error
		utils.LogObj.Infoln("err", err)
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
			Error:   err.Error(),
		}
	}

	return serializer.Response{
		Status:  code,
		Message: e.GetMsg(code),
		Data:    serializer.BuildCart(carts, product, boss),
	}
}

// Update 更新购物车
func (service *CartService) Update(ctx context.Context, uid uint, id string) serializer.Response {
	code := e.Success
	cartDao := dao.NewCartDao(ctx)
	cId, _ := strconv.Atoi(id)

	cart, err := cartDao.ShowCart(uid, uint(cId))
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
			Error:   err.Error(),
		}
	}
	if service.Num != 0 {
		cart.Num = service.Num
	}
	err = cartDao.UpdateCart(uid, uint(cId), cart)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
			Error:   err.Error(),
		}
	}
	return serializer.Response{
		Status:  code,
		Message: e.GetMsg(code),
	}
}

// Delete 移出购物车
func (service *CartService) Delete(ctx context.Context, uid uint, id string) serializer.Response {
	code := e.Success
	cartDao := dao.NewCartDao(ctx)
	cId, _ := strconv.Atoi(id)
	err := cartDao.DeleteCart(uid, uint(cId))
	if err != nil {
		code = e.Error
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

// DeleteAll 清空购物车
func (service *CartService) DeleteAll(ctx context.Context, uid uint) serializer.Response {
	code := e.Success
	cartDao := dao.NewCartDao(ctx)
	err := cartDao.DeleteAllCarts(uid)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
			Error:   err.Error(),
		}
	}
	return serializer.Response{
		Status:  code,
		Message: e.GetMsg(code),
	}
}
