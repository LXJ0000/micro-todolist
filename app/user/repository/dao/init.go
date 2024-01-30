package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"micro-todolist/app/user/repository/model"
	"micro-todolist/conf"
	"strconv"
)

var _db *gorm.DB

func Init() {
	m := conf.Conf.MySQLConfig
	dsn := m.User + ":" + m.Password + "@tcp(" + m.Host + ":" + strconv.Itoa(m.Port) + ")/" + m.DB + "?charset=utf8mb4&parseTime=True&loc=Local"

	//todo ormLogger
	ormLogger := logger.Default.LogMode(logger.Info)
	var err error
	_db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,   //string类型字段默认长度
		DisableDatetimePrecision:  true,  //禁用datetime精度，兼容mysql5.6
		DontSupportRenameColumn:   true,  //不支持重命名索引
		DontSupportRenameIndex:    true,  //不支持重命名字段
		SkipInitializeWithVersion: false, //根据版本自动配置
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: ormLogger,
	})
	if err != nil {
		panic(fmt.Errorf("mysql 连接失败%s", err.Error()))
	}
	sqlDB, _ := _db.DB()
	sqlDB.SetMaxIdleConns(m.MaxIdleConns)
	sqlDB.SetMaxOpenConns(m.MaxOpenConns)

	migration()
}

func migration() {
	if err := _db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&model.User{},
	); err != nil {
		panic(err)
	}
}
