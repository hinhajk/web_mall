package service

import (
	"context"
	"web_mall/dao"
	"web_mall/pkg/e"
	"web_mall/pkg/utils"
	"web_mall/serializer"
)

type Carousel struct {
}

func (service *Carousel) List(context context.Context) serializer.Response {
	carouselDao := dao.NewNCarouselDao(context)
	code := e.Success
	carousel, err := carouselDao.FindCarouselById()
	if err != nil {
		utils.LogObj.Infoln("err", err)
		code = e.Error
		return serializer.Response{
			Status:  code,
			Message: err.Error(),
			Error:   err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildListCarousel(carousel), uint(len(carousel)))

}
