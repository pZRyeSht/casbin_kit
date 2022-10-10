package dao

import (
	"casbin_kit/internal/model"
)

func (d *Dao) CasbinCreate(roleId string, path, method string) error {
	cm := model.CasbinModel{
		PType:  "p",
		RoleId: roleId,
		Path:   path,
		Method: method,
	}
	return cm.Create(d.engine)
}

func (d *Dao) CasbinList(roleID string) [][]string {
	cm := model.CasbinModel{RoleId: roleID}
	return cm.List(d.engine)
}
