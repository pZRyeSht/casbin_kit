package v1

import (
	"fmt"
	kitmodel "github.com/EscAlice/casbin_kit/internal/model"
	"github.com/EscAlice/casbin_kit/internal/service"
	"github.com/EscAlice/casbin_kit/pkg"
	"github.com/gin-gonic/gin"
	"log"
)

type Casbin struct{}

func NewCasbin() Casbin {
	return Casbin{}
}

// Create godoc
// @Summary 新增权限
// @Description 新增权限
// @Tags 权限管理
// @Produce json
// @Security ApiKeyAuth
// @Param body body kitmodel.CasbinCreateRequest true "body"
// @Success 200 {object} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/casbin [post]
func (c Casbin) Create(ctx *gin.Context) {
	var param kitmodel.CasbinCreateRequest
	if err := ctx.ShouldBindJSON(&param); err != nil {
		log.Printf("ShouldBindJSON errs: %v", err)
		pkg.FailWithMessage(fmt.Sprintf("参数解析失败 %s", err.Error()), ctx)
		return
	}

	// 进行插入操作
	svc := service.CasbinServiceApp
	fmt.Println(param)
	if err := svc.CreateCasbin(param.RoleId, param.CasbinInfos); err != nil {
		log.Printf("svc.CasbinCreate err: %v", err)
		pkg.FailWithMessage(fmt.Sprintf("权限验证不通过 %s", err.Error()), ctx)
		return
	}
	pkg.OkWithMessage("权限增加成功", ctx)
	return
}

// List godoc
// @Summary 获取权限列表
// @Produce json
// @Tags 权限管理
// @Security ApiKeyAuth
// @Param data body kitmodel.CasbinListRequest true "角色ID"
// @Success 200 {object} service.CasbinListResponse "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/casbin/list [post]
func (c Casbin) List(ctx *gin.Context) {
	param := kitmodel.CasbinListRequest{}
	if err := ctx.ShouldBindJSON(&param); err != nil {
		log.Printf("ShouldBindJSON errs: %v", err)
		pkg.FailWithMessage(fmt.Sprintf("参数解析失败 %s", err.Error()), ctx)
		return
	}
	// 业务逻辑处理
	svc := service.CasbinServiceApp
	casbins := svc.GetPolicyPathByAuthorityId(param.RoleID)
	pkg.OkWithDetailed(casbins, "权限列表获取成功", ctx)
	return
}
