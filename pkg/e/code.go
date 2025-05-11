package e

// 存放错误状态码信息
const (
	Success      = 200
	Error        = 500
	InvalidParas = 400

	// user模块错误
	ErrorExistUser      = 3001 //用户已存在
	ErrorNotExistUser   = 3002
	ErrorFailEncryption = 3003
	ErrorPwd            = 3004 //	密码加密失败
	ErrorAuthToken      = 3005
	ErrorTokenExpired   = 3006
	ErrorUploadFail     = 3007
	ErrorSendEmailFail  = 3008

	// product 模块错误
	ErrorProductImgUpload = 4001

	// 收藏夹错误
	ErrorFavoritesExist = 5001
)
