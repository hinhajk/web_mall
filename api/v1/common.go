package api

import (
	"encoding/json"
	"errors"
	"web_mall/serializer"
)

func ErrorResponse(err error) serializer.Response {
	var unmarshalTypeError *json.UnmarshalTypeError
	if errors.As(err, &unmarshalTypeError) {
		return serializer.Response{
			Status:  400,
			Message: "JSON类型不匹配",
			Error:   err.Error(),
		}
	}
	return serializer.Response{
		Status:  400,
		Message: "参数错误",
		Error:   err.Error(),
	}
}
