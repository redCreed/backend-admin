package service

import (
	"back-admin/internal/app/common"
	"back-admin/internal/app/sys/dto"
	"back-admin/pkg/xerr"
	"github.com/gin-gonic/gin"
	"strconv"
)

//RoleList 角色列表
func (u *SysSrv) RoleList(param *dto.RoleListParam) (*dto.RoleListResp, error) {
	data, err := u.db.Sys().GetRoleList(param)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (u *SysSrv) AddRole(c *gin.Context, param *dto.AddRoleParam) error {
	//判断是否重名
	role, err := u.db.Sys().GetRoleByField("role_name", param.RoleName)
	if err != nil {
		_, ok := err.(*xerr.CodeError)
		if !ok {
			return err
		}
	}
	rules := make([][]string, 0)
	//校验菜单数据
	if len(param.MenuIds) > 0 {
		menu, err := u.db.Sys().GetMenuByIds(param.MenuIds)
		if err != nil {
			return err
		}
		if len(menu) != len(param.MenuIds) {
			return xerr.NewErrCodeMsg(xerr.ValidateParamErr, "menu_ids参数错误")
		}
		//查询api信息
		api, err := u.db.Sys().GetApiByMenuIds(param.MenuIds)
		if err != nil {
			return err
		}

		for _, v := range api {
			rules = append(rules, []string{strconv.Itoa(role.RoleId), v.Path, v.Method})
		}
	}
	if err := u.db.Sys().AddRole(c.GetInt("user_id"), param); err != nil {
		return err
	}
	//添加权限
	if len(rules) > 0 {
		if _, err := u.enforcer.Enforcer.AddNamedPolicies("p", rules); err != nil {
			return err
		}
	}

	return nil
}

func (u *SysSrv) UpdateRole(c *gin.Context, param *dto.UpdateRoleParam, roleId int) error {
	role, err := u.db.Sys().GetRoleById(roleId)
	if err != nil {
		return err
	}

	roleMenu, err := u.db.Sys().GetMenuByRoleId(roleId)
	if err != nil {
		return err
	}
	menuId := make([]int, 0)
	for _, v := range roleMenu {
		menuId = append(menuId, v.MenuId)
	}
	//管理员
	if role.RoleKey == common.Admin {
		param.MenuIds = make([]int, 0)
		menuId = make([]int, 0)
	}
	param.RoleId = roleId
	if err = u.db.Sys().UpdateRole(c.GetInt("user_id"), param, menuId); err != nil {
		return err
	}

	//查询api
	api, err := u.db.Sys().GetApiByMenuIds(param.MenuIds)
	if err != nil {
		return err
	}
	//删除旧的权限
	if _, err = u.enforcer.DeleteUser(strconv.Itoa(param.RoleId)); err != nil {
		return err
	}
	if len(api) > 0 {
		rules := make([][]string, 0)
		for _, v := range api {
			rules = append(rules, []string{strconv.Itoa(role.RoleId), v.Path, v.Method})
		}
		if _, err := u.enforcer.Enforcer.AddNamedPolicies("p", rules); err != nil {
			return err
		}
	}

	return nil
}

func (u *SysSrv) DeleteRole(c *gin.Context, roleId int) error {
	if err := u.db.Sys().DeleteRole(roleId); err != nil {
		return err
	}
	if _, err := u.enforcer.Enforcer.DeleteRole(strconv.Itoa(roleId)); err != nil {
		return err
	}
	return nil
}
