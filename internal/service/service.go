package service

import (
	"casbin_kit/global"
	"casbin_kit/internal/dao"
	"github.com/gin-gonic/gin"
)

type Service struct {
	ctx *gin.Context
	dao *dao.Dao
}

func NewService(ctx *gin.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.NewDao(global.DBEngine)
	return svc
}


