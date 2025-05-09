package service

import (
	"context"
	"mime/multipart"
	"strconv"
	"sync"
	"web_mall/dao"
	"web_mall/models"
	"web_mall/pkg/e"
	"web_mall/pkg/utils"
	"web_mall/serializer"
)

// ProductService 商品创建、更新、查询接口
type ProductService struct {
	ID            uint   `json:"id" form:"id"`
	ProductName   string `json:"product_name" form:"product_name"`
	Category      uint   `json:"category" form:"category"`
	Title         string `json:"title" form:"title"`
	Info          string `json:"info" form:"info"`
	ImgPath       string `json:"img_path" form:"img_path"`
	Price         string `json:"price" form:"price"`
	OnSale        bool   `json:"onSale" form:"onSale"`
	DisCountPrice string `json:"discount_price" form:"discount_price"`
	Num           int    `json:"num" form:"num"`
	models.BasePage
}

func (service *ProductService) Create(ctx context.Context, uid uint, files []*multipart.FileHeader) serializer.Response {
	var boss *models.User
	var err error
	code := e.Success
	//查找创建的用户
	userDao := dao.NewUserDao(ctx)
	boss, _ = userDao.FindUserById(uid)
	//以第一张作为封面图
	tmp, _ := files[0].Open()
	path, err := UploadProductToLocalStatic(tmp, uid, service.ProductName)
	if err != nil {
		code = e.ErrorProductImgUpload
		utils.LogObj.Infoln(err)
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
			Error:   err.Error(),
		}
	}
	product := &models.Product{
		ProductName:   service.ProductName,
		Category:      service.Category,
		Title:         service.Title,
		Info:          service.Info,
		ImgPath:       path,
		Price:         service.Price,
		OnSale:        true,
		DiscountPrice: service.DisCountPrice,
		Num:           service.Num,
		BossId:        boss.ID,
		BossName:      boss.UserName,
		BossAvatar:    boss.Avatar,
	}
	productDao := dao.NewProductDao(ctx)
	err = productDao.CreateProduct(product)
	if err != nil {
		code = e.Error
		utils.LogObj.Infoln(err)
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
			Error:   err.Error(),
		}
	}
	//并发式创建
	wg := new(sync.WaitGroup)
	wg.Add(len(files))
	for index, file := range files {
		num := strconv.Itoa(index)
		productImgDao := dao.NewProductImageDaoByDB(productDao.DB)
		tmp, _ = file.Open()
		path, err = UploadProductToLocalStatic(tmp, uid, service.ProductName+num)
		if err != nil {
			code = e.ErrorProductImgUpload
			return serializer.Response{
				Status:  code,
				Message: e.GetMsg(code),
				Error:   err.Error(),
			}
		}
		productImg := models.ProductImage{
			ProductId: product.ID,
			ImgPath:   path,
		}
		err = productImgDao.CreateProductImg(&productImg)
		if err != nil {
			code = e.Error
			return serializer.Response{
				Status:  code,
				Message: e.GetMsg(code),
				Error:   err.Error(),
			}
		}
		wg.Done()
	}
	wg.Wait()
	return serializer.Response{
		Status:  code,
		Message: e.GetMsg(code),
		Data:    serializer.BuildProduct(product),
	}
}
