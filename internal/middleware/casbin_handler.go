package middleware

import (
	"casbin_kit/internal/service"
	"casbin_kit/pkg"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

var casbinService = service.CasbinServiceApp

// CasbinHandler 拦截器
func CasbinHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, _ := pkg.GetClaims(ctx)
		// 获取请求的URI
		obj := ctx.Request.URL.Path
		// 获取请求方法
		act := ctx.Request.Method
		// 获取用户的角色
		sub := strconv.Itoa(int(user.RoleId))
		e := casbinService.Casbin()
		fmt.Println(obj, act, sub)
		// 判断策略中是否存在
		success, _ := e.Enforce(sub, obj, act)
		if success {
			log.Println("权限验证通过")
			ctx.Next()
		} else {
			log.Printf("e.Enforce err: %s", "权限不足")
			pkg.FailWithDetailed(gin.H{}, "权限不足", ctx)
			ctx.Abort()
			return
		}
	}
}
