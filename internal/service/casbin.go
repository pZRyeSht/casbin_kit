package service

import (
	"casbin_kit/global"
	kitmodel "casbin_kit/internal/model"
	"errors"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"

	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"sync"
)

type CasbinService struct{}

var CasbinServiceApp = new(CasbinService)

func (casbinService *CasbinService) CreateCasbin(RoleID uint, casbinInfos []kitmodel.CasbinInfo) error {
	authorityId := strconv.Itoa(int(RoleID))
	casbinService.ClearCasbin(0, authorityId)
	var rules [][]string
	for _, v := range casbinInfos {
		rules = append(rules, []string{authorityId, v.Path, v.Method})
		fmt.Println(v)
	}
	e := casbinService.Casbin()
	success, _ := e.AddPolicies(rules)
	if !success {
		return errors.New("存在相同api,添加失败,请联系管理员")
	}
	err := e.InvalidateCache()
	if err != nil {
		return err
	}
	return nil
}

func (casbinService *CasbinService) UpdateCasbinApi(oldPath string, newPath string, oldMethod string, newMethod string) error {
	err := global.DBEngine.Model(&gormadapter.CasbinRule{}).Where("v1 = ? AND v2 = ?", oldPath, oldMethod).Updates(map[string]interface{}{
		"v1": newPath,
		"v2": newMethod,
	}).Error
	e := casbinService.Casbin()
	err = e.InvalidateCache()
	if err != nil {
		return err
	}
	return err
}

func (casbinService *CasbinService) GetPolicyPathByAuthorityId(RoleID uint) (pathMaps []kitmodel.CasbinInfo) {
	e := casbinService.Casbin()
	roleId := strconv.Itoa(int(RoleID))
	list := e.GetFilteredPolicy(0, roleId)
	for _, v := range list {
		pathMaps = append(pathMaps, kitmodel.CasbinInfo{
			Path:   v[1],
			Method: v[2],
		})
	}
	return pathMaps
}

func (casbinService *CasbinService) ClearCasbin(v int, p ...string) bool {
	e := casbinService.Casbin()
	success, _ := e.RemoveFilteredPolicy(v, p...)
	return success
}

var (
	cachedEnforcer *casbin.CachedEnforcer
	once           sync.Once
)

func (casbinService *CasbinService) Casbin() *casbin.CachedEnforcer {
	once.Do(func() {
		a, _ := gormadapter.NewAdapterByDB(global.DBEngine)
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
			fmt.Println("字符串加载模型失败!", err)
			return
		}
		cachedEnforcer, _ = casbin.NewCachedEnforcer(m, a)
		cachedEnforcer.SetExpireTime(60 * 60)
		_ = cachedEnforcer.LoadPolicy()
	})
	return cachedEnforcer
}
