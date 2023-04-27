package data

import (
	"fmt"
	"gin-jj/app"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func InitMysql() {
	logConfig := logger.New(
		nil,
		logger.Config{
			SlowThreshold:             0,                  // 慢 SQL 阈值
			LogLevel:                  logger.LogLevel(4), // 日志级别
			IgnoreRecordNotFoundError: false,              // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,              // 是否启用彩色打印
		},
	)

	configs := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: app.C.Mysql.Prefix,
		},
		Logger: logConfig,
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		app.C.Mysql.Username,
		app.C.Mysql.Password,
		app.C.Mysql.Host,
		app.C.Mysql.Port,
		app.C.Mysql.Database,
		app.C.Mysql.Charset,
	)

	if db, err := gorm.Open(mysql.Open(dsn), configs); err != nil {
		panic("Mysql connection failed：" + err.Error())
	} else {
		sqlDB, _ := db.DB()
		// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
		sqlDB.SetMaxIdleConns(app.C.Mysql.MaxIdleConns)
		// SetMaxOpenConns 设置打开数据库连接的最大数量。
		sqlDB.SetMaxOpenConns(app.C.Mysql.MaxOpenConns)
		// SetConnMaxLifetime 设置了连接可复用的最大时间。
		sqlDB.SetConnMaxLifetime(app.C.Mysql.MaxLifetime)
		app.Db = db
	}
}
