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
	OnSale        bool   `json:"onSale" form:"onSale" default:"false"`
	DisCountPrice string `json:"discount_price" form:"discount_price"`
	Num           int    `json:"num" form:"num"`
	models.BasePage
}

// Create 创建商品
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
		OnSale:        false,
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

// ProductList 获取商品列表
func (service *ProductService) ProductList(ctx context.Context) serializer.Response {
	var products []*models.Product
	var err error
	code := e.Success
	if service.PageSize == 0 {
		service.PageSize = 15
	} //默认页数为15

	//查询某类商品
	condition := make(map[string]interface{})
	if service.Category != 0 {
		condition["category_id"] = service.Category
	}
	productDao := dao.NewProductDao(ctx)
	total, err := productDao.CountProductByCategory(condition)
	if err != nil {
		code = e.Error
		utils.LogObj.Infoln(err)
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
			Error:   err.Error(),
		}
	}

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		productDao = dao.NewProductDaoByDB(productDao.DB)
		products, _ = productDao.ListCountProductByCategory(condition, service.BasePage)
		wg.Done()
	}()
	wg.Wait()
	return serializer.BuildListResponse(serializer.BuildProducts(products), uint(total))
}

// Search 搜索商品
func (service *ProductService) Search(ctx context.Context, uid uint) serializer.Response {
	code := e.Success
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	productDao := dao.NewProductDao(ctx)
	products, count, err := productDao.SearchProduct(service.Info, service.BasePage)
	if err != nil {
		code = e.Error
		utils.LogObj.Infoln(err)
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
			Error:   err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildProducts(products), uint(count))
}

// Show 显示商品详细信息和图片信息
func (service *ProductService) Show(ctx context.Context, id string) serializer.Response {
	code := e.Success
	productDao := dao.NewProductDao(ctx)
	pid, _ := strconv.Atoi(id)
	product, err := productDao.ShowProductByID(uint(pid))
	if err != nil {
		code = e.Error
		utils.LogObj.Infoln(err)
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
			Error:   err.Error(),
		}
	}
	return serializer.Response{
		Status:  code,
		Data:    serializer.BuildProduct(product),
		Message: e.GetMsg(code),
	}
}

// Delete 删除商品
func (service *ProductService) Delete(ctx context.Context, uid uint, id string) serializer.Response {
	productDao := dao.NewProductDao(ctx)
	pid, _ := strconv.Atoi(id)
	code := e.Success
	err := productDao.DeleteProduct(uid, uint(pid))
	if err != nil {
		code = e.Error
		utils.LogObj.Infoln(err)
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

// Update 更新商品信息
func (service *ProductService) Update(ctx context.Context, uid uint, id string) serializer.Response {
	code := e.Success
	productDao := dao.NewProductDao(ctx)
	pid, _ := strconv.Atoi(id)
	product, err := productDao.ShowProductByID(uint(pid))
	product.OnSale = service.OnSale
	if service.DisCountPrice != "" {
		product.DiscountPrice = service.DisCountPrice
	}
	if service.Info != "" {
		product.Info = service.Info
	}
	if service.Category != 0 {
		product.Category = service.Category
	}
	if service.Title != "" {
		product.Title = service.Title
	}
	if service.ProductName != "" {
		product.ProductName = service.ProductName
	}
	if service.Num != 0 {
		product.Num = service.Num
	}
	if service.Price != "" {
		product.Price = service.Price
	}
	err = productDao.UpdateById(uid, uint(pid), product)
	if err != nil {
		code = e.Error
		utils.LogObj.Infoln(err)
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
			Error:   err.Error(),
		}
	}
	return serializer.Response{
		Status:  code,
		Message: e.GetMsg(code),
		Data:    serializer.BuildProduct(product),
	}
}

// UpdateAvatar 更新商品图片信息
func (service *ProductService) UpdateAvatar(ctx context.Context, uid uint, id string, files []*multipart.FileHeader) serializer.Response {
	code := e.Success
	productDao := dao.NewProductDao(ctx)
	pid, _ := strconv.Atoi(id)
	product, err := productDao.ShowProductByID(uint(pid))
	if err != nil {
		code = e.Error
		utils.LogObj.Infoln(err)
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
			Error:   err.Error(),
		}
	}
	tmp, _ := files[0].Open()
	if tmp != nil {
		path, err1 := UploadProductToLocalStatic(tmp, uid, product.ProductName)
		if err1 != nil {
			utils.LogObj.Infoln(err1)
			return serializer.Response{
				Status:  code,
				Message: e.GetMsg(code),
				Error:   err1.Error(),
			}
		}
		product.ImgPath = path
	}
	err = productDao.UpdateById(uid, uint(pid), product)
	if err != nil {
		code = e.Error
		utils.LogObj.Infoln(err)
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
			Error:   err.Error(),
		}
	}
	return serializer.Response{
		Status:  code,
		Message: e.GetMsg(code),
		Data:    serializer.BuildProduct(product),
	}
}
