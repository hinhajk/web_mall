package dao

// 该文件中进行数据库连接

import (
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
	"log"
	"time"
)

// 设置数据库连接池
func setPool(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Println(err)
		return
	}
	sqlDB.SetMaxIdleConns(20)           // 设置最大连接数
	sqlDB.SetConnMaxLifetime(time.Hour) //设置最长活跃时间
	sqlDB.SetMaxOpenConns(100)
}

var _db *gorm.DB

func DataBaseSparate(connRead, coonWrite string) {
	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       connRead,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true, //禁止DateTime精度
		DontSupportRenameIndex:    true, //重命名索引的话，要先将索引删掉
		DontSupportRenameColumn:   true, //用change重命名列
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //命名时单数化，不加s
		},
	})
	if err != nil {
		log.Println(err)
		return
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetMaxOpenConns(100)
	_db = db

	//主从配置
	_ = _db.Use(dbresolver.Register(dbresolver.Config{
		Sources:  []gorm.Dialector{mysql.Open(coonWrite)},
		Replicas: []gorm.Dialector{mysql.Open(connRead), mysql.Open(connRead)},
		Policy:   dbresolver.RandomPolicy{},
	}))
	MigrationSeparate()
}

func NewDBClient(ctx context.Context) *gorm.DB {
	db := _db
	return db.WithContext(ctx)
}
