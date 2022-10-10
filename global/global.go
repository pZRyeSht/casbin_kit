package global

import (
	"casbin_kit/config"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var (
	SERVERCONF config.Conf
	DBEngine *gorm.DB
	CONF *viper.Viper
)