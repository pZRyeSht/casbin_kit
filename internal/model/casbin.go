package model

import (
	"casbin_kit/global"
	"errors"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.uber.org/zap"
	"sync"
)

type CasbinModel struct {
	PType  string `json:"p_type" gorm:"column:p_type" description:"策略类型"`
	RoleId string `json:"role_id" gorm:"column:v0" description:"角色ID"`
	Path   string `json:"path" gorm:"column:v1" description:"api路径"`
	Method string `json:"method" gorm:"column:v2" description:"访问方法"`
}

func (c *CasbinModel) TableName() string {
	return "casbin_rule"
}

func (c *CasbinModel) Create(db *gorm.DB) error {
	e := Casbin()
	if success, _ := e.AddPolicy(c.RoleId, c.Path, c.Method); success == false {
		return errors.New("存在相同的API，添加失败")
	}
	return nil
}

func (c *CasbinModel) Update(db *gorm.DB, values interface{}) error {
	if err := db.Model(c).Where("v1 = ? AND v2 = ?", c.Path, c.Method).Update(values).Error; err != nil {
		return err
	}
	return nil
}

func (c *CasbinModel) List(db *gorm.DB) [][]string {
	e := Casbin()
	policy := e.GetFilteredPolicy(0, c.RoleId)
	return policy
}

var (
	cachedEnforcer *casbin.CachedEnforcer
	once           sync.Once
)

// Casbin
// @function: Casbin
// @description: 持久化到数据库  引入自定义规则
// @return: *casbin.Enforcer
func Casbin() *casbin.CachedEnforcer {
	s := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		global.SERVERCONF.Database.Username,
		global.SERVERCONF.Database.Password,
		global.SERVERCONF.Database.Host,
		global.SERVERCONF.Database.DBName,
		global.SERVERCONF.Database.Charset,
		global.SERVERCONF.Database.ParseTime,
	)
	db, _ := gorm.Open(global.SERVERCONF.Database.DBType, s)
	once.Do(func() {
		adapter := gormadapter.NewAdapterByDB(db)
		text := `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act
		`
		m, err := model.NewModelFromString(text)
		if err != nil {
			zap.L().Error("字符串加载模型失败!", zap.Error(err))
			return
		}
		cachedEnforcer, _ = casbin.NewCachedEnforcer(m, adapter)
		cachedEnforcer.SetExpireTime(60 * 60)
		_ = cachedEnforcer.LoadPolicy()
	})
	return cachedEnforcer
}

// ClearCasbin
// @function: ClearCasbin
// @description: 清除匹配的权限
// @param: v int, p ...string
// @return: bool
func ClearCasbin(v int, p ...string) bool {
	e := Casbin()
	success, _ := e.RemoveFilteredPolicy(v, p...)
	return success
}
