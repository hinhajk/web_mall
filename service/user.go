package service

import (
	"context"
	"gopkg.in/mail.v2"
	"mime/multipart"
	"time"
	"web_mall/config"
	"web_mall/dao"
	"web_mall/models"
	"web_mall/pkg/e"
	"web_mall/pkg/utils"
	"web_mall/serializer"
)

// UserService 用户注册、登录、修改、头像上传服务接口
type UserService struct {
	NickName string `json:"nick_name" form:"nick_name"`
	UserName string `json:"user_name" form:"user_name"`
	PassWord string `json:"password" form:"password"`
	Gender   string `json:"gender" form:"gender"`
	Key      string `json:"key" form:"key"` //加密密钥（一般情况下前端、后端都要验证一次；这里后端不做验证，在前端验证）
}

// SendEmailService 发送邮箱
type SendEmailService struct {
	Email         string `json:"email" form:"email"`
	PassWord      string `json:"password" form:"password"`
	OperationType uint   `json:"operation_type" form:"operation_type"` //1.绑定邮箱；2.解绑邮箱；3.改密码
}

// ValidEmailService 验证邮箱
type ValidEmailService struct {
}

type ShowMoneyService struct {
	Key string `json:"key" form:"key"`
}

// Register 用户注册
func (service *UserService) Register(ctx context.Context) serializer.Response {
	var user models.User
	code := e.Success
	if service.Key == "" || len(service.Key) != 16 {
		code = e.Error
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
			Error:   "密钥长度不足",
		}
	} //如果密钥为空或者长度不符合要求，返回错误信息

	//密钥加密功能：初始金额要进行密文存储
	//10000--->对称加密操作
	utils.Encrypt.SetKey(service.Key)

	userDao := dao.NewUserDao(ctx)
	_, exist, err := userDao.ExistOrNotByUserName(service.UserName)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
		}
	}
	if !exist {
		code = e.ErrorExistUser
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
		}
	}
	user = models.User{
		UserName: service.UserName,
		Nickname: service.NickName,
		Status:   models.Active,
		Avatar:   "avatar.JPG",
		Gender:   service.Gender,
		Money:    utils.Encrypt.AesEncoding("10000"), //初始金额加密
	}
	//密码加密
	if err = user.SetPassWord(service.PassWord); err != nil {
		code = e.ErrorFailEncryption
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
		}
	}
	//创建用户
	err = userDao.CreateUser(&user)
	if err != nil {
		code = e.Error
	}
	return serializer.Response{
		Status:  code,
		Message: e.GetMsg(code),
		Data:    serializer.BuildUser(&user),
	}
}

// Login 用户登录
func (service *UserService) Login(ctx context.Context) serializer.Response {
	var user *models.User
	code := e.Success
	userDao := dao.NewUserDao(ctx)
	user, exist, err := userDao.ExistOrNotByUserName(service.UserName)
	if exist == false || err != nil {
		code = e.ErrorNotExistUser
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
			Data:    "用户不存在，请先注册",
		}
	}

	//检查密码
	if user.CheckPassWord(service.PassWord) == false {
		code = e.ErrorPwd
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
			Data:    "密码错误， 请重新输入",
		}
	}

	//token签发
	//http是无状态的，因此服务器不知道访问者是谁，；让请求携带cookie值，服务器就可以知道访问者是谁
	token, err := utils.GenerateToken(user.ID, service.UserName, 0)
	if err != nil {
		code = e.ErrorAuthToken
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status:  code,
		Message: e.GetMsg(code),
		Data:    serializer.TokenData{User: serializer.BuildUser(user), Token: token},
	}
}

// Update 用户信息修改
func (service *UserService) Update(ctx context.Context, uid uint) serializer.Response {
	var user *models.User
	var err error
	code := e.Success
	//找到这个用户
	userDao := dao.NewUserDao(ctx)
	user, err = userDao.FindUserById(uid)
	//修改昵称nickname
	if service.NickName != "" {
		user.Nickname = service.NickName
	}
	if service.Gender != "" {
		user.Gender = service.Gender
	}
	err = userDao.UpdateById(uid, user)
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
		Data:    serializer.BuildUser(user),
	}
}

// LoadAvatar 上传头像
func (service *UserService) LoadAvatar(ctx context.Context, uid uint, file multipart.File, filesSize int64) serializer.Response {
	var user *models.User
	var err error
	code := e.Success
	userDao := dao.NewUserDao(ctx)
	user, err = userDao.FindUserById(uid)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
			Error:   err.Error(),
		}
	}
	//保存图片到本地的函数
	path, err := UploadAvatarToLocalAvatar(file, uid, user.UserName)
	if err != nil {
		code = e.ErrorUploadFail
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
			Error:   err.Error(),
		}
	}
	user.Avatar = path
	err = userDao.UpdateById(uid, user)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status:  code,
		Message: e.GetMsg(code),
		Data:    serializer.BuildUser(user),
	}
}

// Send 发送邮箱
func (service *SendEmailService) Send(ctx context.Context, uid uint) serializer.Response {
	code := e.Success
	var address string        //发送地址
	var notice *models.Notice //邮箱模板通知
	//每发一次邮箱进行一次验证，确保是该次验证，所以需要签发token
	token, err := utils.GenerateEmailToken(uid, service.OperationType, service.Email, service.PassWord)
	if err != nil {
		code = e.ErrorAuthToken
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
			//Message: "token签发错误",
			Error: err.Error(),
		}
	}

	noticeDao := dao.NewNoticeDao(ctx)
	notice, err = noticeDao.FindNoticeById(service.OperationType)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
			Error:   err.Error(),
		}
	}
	address = config.ValidEmail + token //发送方
	m := mail.NewMessage()
	m.SetHeader("From", config.SmtpEmail)
	m.SetHeader("To", service.Email)
	m.SetHeader("Subject", "1111")
	m.SetBody("text/html", notice.Text+address)
	d := mail.NewDialer(config.SmtpHost, 465, config.SmtpEmail, config.SmtpPass)
	d.StartTLSPolicy = mail.MandatoryStartTLS
	if err := d.DialAndSend(m); err != nil {
		code = e.ErrorSendEmailFail
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

// Valid 验证邮箱
func (service *ValidEmailService) Valid(ctx context.Context, token string) serializer.Response {
	var userId uint
	var email string
	var password string
	var operationType uint
	code := e.Success
	//验证token
	if token == "" {
		code = e.InvalidParas
	} else {
		claims, err := utils.ParseEmailToken(token)
		if err != nil {
			code = e.ErrorAuthToken
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = e.ErrorTokenExpired
		} else {
			userId = claims.UserID
			email = claims.Email
			password = claims.Password
			operationType = claims.OperationType
		}
	}
	if code != e.Success {
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
		}
	}

	//获取用户信息
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.FindUserById(userId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
			Error:   err.Error(),
		}
	}
	if operationType == 1 {
		user.Email = email
	} else if operationType == 2 {
		//解绑邮箱
		user.Email = ""
	} else if operationType == 3 {
		//修改密码
		err = user.SetPassWord(password)
		if err != nil {
			code = e.Error
			return serializer.Response{
				Status:  code,
				Message: e.GetMsg(code),
			}
		}
	}
	err = userDao.UpdateById(userId, user)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status:  code,
		Message: e.GetMsg(code),
		Data:    serializer.BuildUser(user),
	}
}

// Show 显示用户金额
func (service *ShowMoneyService) Show(ctx context.Context, uid uint) serializer.Response {
	code := e.Success
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.FindUserById(uid)
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
		Data:    serializer.BuildMoney(user, service.Key),
	}
}
