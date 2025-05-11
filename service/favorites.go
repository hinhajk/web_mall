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

type FavoritesService struct {
	ProductId  uint `json:"product_id" form:"product_id"`
	BossId     uint `json:"boss_id" form:"boss_id"`
	FavoriteId uint `json:"favorite_id" form:"favorite_id"`
	models.BasePage
}

// Delete 删除收藏夹
func (service *FavoritesService) Delete(ctx context.Context, uid uint, id string) serializer.Response {
	code := e.Success
	favoritesDao := dao.NewFavoritesDao(ctx)
	fId, _ := strconv.Atoi(id)
	err := favoritesDao.DeleteFavorites(uint(fId), uid)
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

// Create 创建收藏夹
func (service *FavoritesService) Create(ctx context.Context, uid uint) serializer.Response {
	code := e.Success
	favoriteDao := dao.NewFavoritesDao(ctx)
	exist, err := favoriteDao.FavoriteExistOrNot(service.ProductId, uid)
	if exist {
		code = e.ErrorFavoritesExist
		utils.LogObj.Infoln("err", err)
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
			Error:   err.Error(),
		}
	}
	//查找用户
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.FindUserById(uid)
	if err != nil {
		code = e.Error
		utils.LogObj.Infoln("err", err)
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
			Error:   err.Error(),
		}
	}
	//查找老板
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

	//查找商品
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

	favorite := &models.Favorite{
		User:      *user,
		Boss:      *boss,
		Product:   *product,
		UserId:    uid,
		BossId:    service.BossId,
		ProductId: service.ProductId,
	}

	//创建收藏夹
	err = favoriteDao.CreateFavorites(favorite)
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
	}
}

// Show 查询收藏夹
func (service *FavoritesService) Show(ctx context.Context, uid uint) serializer.Response {
	code := e.Success
	favoritesDao := dao.NewFavoritesDao(ctx)
	favorites, err := favoritesDao.ListFavoritesByUid(uid)
	if err != nil {
		utils.LogObj.Infoln("err", err)
		code = e.Error
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
			Error:   err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildFavorites(ctx, favorites), uint(len(favorites)))
}
