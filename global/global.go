package global

import (
	"casbin_kit/config"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	SERVERCONF config.Conf
	DBEngine   *gorm.DB
	CONF       *viper.Viper
)
