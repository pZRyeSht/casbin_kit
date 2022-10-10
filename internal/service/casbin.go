package service

import (
	kitmodel "casbin_kit/internal/model"
)

func (s Service) CasbinCreate(param *kitmodel.CasbinCreateRequest) error {
	for _, v := range param.CasbinInfos {
		err := s.dao.CasbinCreate(param.RoleId, v.Path, v.Method)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s Service) CasbinList(param *kitmodel.CasbinListRequest) [][]string {
	return s.dao.CasbinList(param.RoleID)
}
