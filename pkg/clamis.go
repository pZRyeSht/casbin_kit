package pkg

import (
	"casbin_kit/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func GetClaims(c *gin.Context) (*model.CustomClaims, error) {
	//token := c.Request.Header.Get("x-token")
	//j := NewJWT()
	//claims, err := j.ParseToken(token)
	//if err != nil {
	//	return nil, errors.New("从Gin的Context中获取从jwt解析信息失败, 请检查请求头是否存在x-token且claims是否为规定结构")
	//}
	// todo jwt token
	return &model.CustomClaims{
		BaseClaims: model.BaseClaims{
			ID:     1,
			AuthId: "admin",
		},
		BufferTime:     0,
		StandardClaims: jwt.StandardClaims{},
	}, nil
}
