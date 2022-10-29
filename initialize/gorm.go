package initialize

import (
	"fmt"
	"github.com/EscAlice/casbin_kit/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GormMysql() *gorm.DB {
	m := global.SERVERCONF.Database
	if m.DBName == "" {
		return nil
	}
	s := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		global.SERVERCONF.Database.Username,
		global.SERVERCONF.Database.Password,
		global.SERVERCONF.Database.Host,
		global.SERVERCONF.Database.DBName,
		global.SERVERCONF.Database.Charset,
		global.SERVERCONF.Database.ParseTime,
	)
	mysqlConfig := mysql.Config{
		DSN:                       s,     // DSN data source name
		DefaultStringSize:         255,   // string 类型字段的默认长度
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig)); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleTime)
		sqlDB.SetMaxOpenConns(m.MaxOpenConn)
		return db
	}
}
