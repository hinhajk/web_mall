package service

import (
	"context"
	"web_mall/dao"
	"web_mall/pkg/e"
	"web_mall/pkg/utils"
	"web_mall/serializer"
)

type CategoryService struct {
}

func (service *CategoryService) ListCategory(ctx context.Context) serializer.Response {
	categoryDao := dao.NewCategoryDao(ctx)
	code := e.Success
	category, err := categoryDao.FindCategoryById()
	if err != nil {
		utils.LogObj.Infoln("err", err)
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
			Error:   err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildListCategories(category), uint(len(category)))
}
