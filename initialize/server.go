package initialize

import (
	"casbin_kit/internal/router"
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	Router := router.NewRouter()
	port := ":8888"
	s := initServer(port, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	fmt.Println("server run success on", port)

	fmt.Printf(`
	欢迎使用 kit_casbin
	当前版本:v1.0.0
	默认前端文件运行地址:http://127.0.0.1%s
`, port)
	if err := s.ListenAndServe(); err != nil {
		fmt.Println("ListenAndServe err ", err.Error())
	}
}

func initServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 20 * time.Second
	s.WriteTimeout = 20 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}
