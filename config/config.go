package config

import (
	"fmt"
	"gopkg.in/ini.v1"
	"strings"
	"web_mall/dao"
)

// 将config.ini中的配置文件读取出来
var (
	AppMode  string
	HttpPort string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	RedisDb     string
	RedisAddr   string
	RedisPw     string
	RedisDbName string

	ValidEmail string
	SmtpHost   string
	SmtpEmail  string
	SmtpPass   string

	Host        string
	ProductPath string
	AvatarPath  string
)

// Init 初始化函数
// 利用init依赖读取ini文件内容传给file变量
func Init() {
	file, err := ini.Load("./config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径")
	}
	LoadServer(file)
	LoadMysql(file)
	LoadRedis(file)
	LoadEmail(file)
	LoadPhotoPath(file)

	// MySQL读写分离实现
	//mysql读
	pathRead := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8&parseTime=True"}, "")
	//mysql写
	pathWrite := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8&parseTime=True"}, "")
	dao.DataBaseSparate(pathRead, pathWrite)
}

// LoadServer 将服务器配置信息给加载出来
func LoadServer(file *ini.File) {
	AppMode = file.Section("service").Key("AppMode").String()
	HttpPort = file.Section("service").Key("HttpPort").String()
	// 将file 文件中server模块中键值为"AppMode"或者其他的部分读取出来兵器转换为对应的string类型
}

// LoadMysql 将MySQL配置信息加载出来
func LoadMysql(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
	DbName = file.Section("mysql").Key("DbName").String()
}

func LoadRedis(file *ini.File) {
	RedisAddr = file.Section("redis").Key("RedisAddr").String()
	RedisDb = file.Section("redis").Key("RedisDb").String()
	RedisDbName = file.Section("redis").Key("RedisDbName").String()
	RedisPw = file.Section("redis").Key("RedisPw").String()
}

func LoadEmail(file *ini.File) {
	SmtpHost = file.Section("email").Key("SmtpHost").String()
	SmtpEmail = file.Section("email").Key("SmtpEmail").String()
	SmtpPass = file.Section("email").Key("SmtpPass").String()
	ValidEmail = file.Section("email").Key("ValidEmail").String()
}

func LoadPhotoPath(file *ini.File) {
	AvatarPath = file.Section("path").Key("AvatarPath").String()
	Host = file.Section("path").Key("Host").String()
	ProductPath = file.Section("path").Key("ProductPath").String()
}
