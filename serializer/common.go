package serializer

import "web_mall/pkg/e"

type Response struct {
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Error   string      `json:"e"`
}

type TokenData struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}

type DataList struct {
	Item  interface{} `json:"item"`
	Total uint        `json:"total"`
}

func BuildListResponse(data interface{}, total uint) Response {
	return Response{
		Status: e.Success,
		Data: DataList{
			Item:  data,
			Total: total,
		},
		Message: "ok",
	}
}
