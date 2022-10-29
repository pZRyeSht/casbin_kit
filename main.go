package main

import (
	"github.com/EscAlice/casbin_kit/global"
	"github.com/EscAlice/casbin_kit/initialize"
)

// @title kit casbin
// @version 1.0
// @description the casbin kit sample code
// @termsOfService https://github.com/EscAlice
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name token
// @BasePath /
func main() {
	// 解析配置文件
	global.CONF = initialize.Viper()
	// 初始化数据库
	global.DBEngine = initialize.GormMysql()
	// 初始化路由
	initialize.RunWindowsServer()
}
