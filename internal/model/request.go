package model

import (
	"github.com/golang-jwt/jwt/v4"
)

// CustomClaims Custom claims structure
type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.StandardClaims
}

type BaseClaims struct {
	ID     uint
	RoleId uint
}

type CasbinInfo struct {
	Path   string `json:"path" form:"path"`
	Method string `json:"method" form:"method"`
}

type CasbinCreateRequest struct {
	RoleId      uint         `json:"role_id" form:"role_id" description:"角色ID"`
	CasbinInfos []CasbinInfo `json:"casbin_infos" description:"权限模型列表"`
}

type CasbinListRequest struct {
	RoleID uint `json:"role_id" form:"role_id"`
}
