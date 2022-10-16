package router

import (
	v1 "casbin_kit/internal/api/v1"
	"casbin_kit/internal/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	casbin := v1.NewCasbin()
	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleware.CasbinHandler())
	{
		// 测试路由
		apiv1.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})

		// 权限策略管理
		apiv1.POST("/casbin", casbin.Create)
		apiv1.POST("/casbin/list", casbin.List)
	}
	return r
}
