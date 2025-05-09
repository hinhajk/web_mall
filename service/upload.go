package service

import (
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"web_mall/config"
)

func UploadAvatarToLocalAvatar(file multipart.File, uid uint, userName string) (filePath string, err error) {
	bId := strconv.Itoa(int(uid)) // 路径拼接
	basePath := "." + config.AvatarPath + "user" + bId + "/"
	if !DirExistOrNot(basePath) {
		CreateDir(basePath)
	}
	avatarPath := basePath + userName + " .jpg"
	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	err = os.WriteFile(avatarPath, content, 0666)
	if err != nil {
		return
	}
	return "user" + bId + "/" + userName + " .jpg", nil
}

func UploadProductToLocalStatic(file multipart.File, uid uint, ProductName string) (filePath string, err error) {
	bId := strconv.Itoa(int(uid)) // 路径拼接
	basePath := "." + config.ProductPath + "boss" + bId + "/"
	if !DirExistOrNot(basePath) {
		CreateDir(basePath)
	}
	productPath := basePath + ProductName + " .jpg"
	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	err = os.WriteFile(productPath, content, 0666)
	if err != nil {
		return
	}
	return "boss" + bId + "/" + ProductName + " .jpg", nil
}

// DirExistOrNot 判断文件夹路径是否存在
func DirExistOrNot(filePath string) bool {
	s, err := os.Stat(filePath)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// CreateDir 创建文件夹
func CreateDir(dirName string) bool {
	err := os.MkdirAll(dirName, 755)
	if err != nil {
		return false
	}
	return true
}
