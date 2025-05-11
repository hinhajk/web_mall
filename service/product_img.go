package service

import (
	"context"
	"strconv"
	"web_mall/dao"
	"web_mall/pkg/e"
	"web_mall/pkg/utils"
	"web_mall/serializer"
)

type ListProductImg struct {
}

func (service *ListProductImg) ListImgs(ctx context.Context, pID string) serializer.Response {
	code := e.Success
	productImgDao := dao.NewProductImageDao(ctx)
	productId, _ := strconv.Atoi(pID)
	productImgs, err := productImgDao.ListProductImg(uint(productId))
	if err != nil {
		code = e.Error
		utils.LogObj.Infoln(err)
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
			Error:   err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildProductImages(productImgs), uint(len(productImgs)))
}
