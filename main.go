package main

import (
	"casbin_kit/global"
	"casbin_kit/internal/core"
	"casbin_kit/pkg"
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
	global.CONF = pkg.Viper()
	core.RunWindowsServer()
}
