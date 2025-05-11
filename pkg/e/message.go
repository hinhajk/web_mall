package e

var MsgFlags = map[int]string{
	Success:               "success",
	Error:                 "failed",
	InvalidParas:          "参数错误",
	ErrorExistUser:        "用户名已存在",
	ErrorFailEncryption:   "密码加密失败",
	ErrorNotExistUser:     "用户名不存在",
	ErrorPwd:              "密码错误",
	ErrorAuthToken:        "token 认证失败",
	ErrorTokenExpired:     "token 已过期",
	ErrorUploadFail:       "上传失败",
	ErrorSendEmailFail:    "邮件发送失败",
	ErrorProductImgUpload: "商品图片上传错误",
	ErrorFavoritesExist:   "商品已收藏",
}

// GetMsg获取状态码对应的信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if !ok {
		return MsgFlags[Error]
	}
	return msg
}
