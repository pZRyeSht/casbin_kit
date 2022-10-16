package initialize

import (
	"casbin_kit/global"
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

// Viper 优先级: 命令行 > 环境变量 > 默认值
func Viper(path ...string) *viper.Viper {
	var config string
	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		if config == "" { // 判断命令行参数是否为空
			if configEnv := os.Getenv(global.ConfigEnv); configEnv == "" { // 判断 internal.ConfigEnv 常量存储的环境变量是否为空
				switch gin.Mode() {
				case gin.DebugMode:
					config = global.ConfigDefaultFile
					fmt.Printf("使用gin模式的%s环境名称启动,config的路径为%s\n", gin.EnvGinMode, global.ConfigDefaultFile)
				case gin.ReleaseMode:
					config = global.ConfigReleaseFile
					fmt.Printf("使用gin模式的%s环境名称启动,config的路径为%s\n", gin.EnvGinMode, global.ConfigReleaseFile)
				case gin.TestMode:
					config = global.ConfigTestFile
					fmt.Printf("使用gin模式的%s环境名称启动,config的路径为%s\n", gin.EnvGinMode, global.ConfigTestFile)
				}
			} else { // internal.ConfigEnv 常量存储的环境变量不为空 将值赋值于config
				config = configEnv
				fmt.Printf("使用%s环境变量启动,config的路径为%s\n", global.ConfigEnv, config)
			}
		} else { // 命令行参数不为空 将值赋值于config
			fmt.Printf("使用命令行的-c参数传递的值启动,config的路径为%s\n", config)
		}
	} else { // 函数传递的可变参数的第一个值赋值于config
		config = path[0]
		fmt.Printf("使用func Viper()传递的值启动,config的路径为%s\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.SERVERCONF); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&global.SERVERCONF); err != nil {
		fmt.Println(err)
	}
	return v
}
