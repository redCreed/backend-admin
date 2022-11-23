package service

import (
	"back-admin/api/models"
	"back-admin/internal/app/sys/dto"
	"back-admin/pkg/xerr"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"strconv"
)

func (u *SysSrv) DeleteMenu(c *gin.Context, menuId int) error {
	ids := []int{menuId}
	menu, err := u.db.Sys().GetMenuByIds(ids)
	if err != nil {
		return err
	}
	if len(menu) == 0 {
		return xerr.NewErrCode(xerr.DataNoExist)
	}

	//获取该菜单api
	api, err := u.db.Sys().GetApiByMenuIds(ids)
	if err != nil {
		return err
	}

	if err = u.db.Sys().DeleteMenu(menuId); err != nil {
		return err
	}

	//删除api权限
	if len(api) > 0 {
		for _, v := range api {
			if _, err = u.enforcer.DeletePermission(v.Path, v.Method); err != nil {
				return err
			}
		}
	}

	return nil
}

func (u *SysSrv) UpdateMenu(c *gin.Context, param *dto.UpdateMenuParam, menuId int) error {
	param.MenuId = menuId
	//判断菜单重复
	m, err := u.db.Sys().GetMenuByField("menu_name", param.MenuName)
	if err != nil {
		return err
	}
	if m.MenuId > 0 {
		return xerr.NewErrCode(xerr.DataHasExist)
	}

	//判断新apis的合法性
	api, err := u.db.Sys().GetApiByIds(param.Apis)
	if err != nil {
		return err
	}
	apiId := make([]int, 0, len(api))
	for _, v := range api {
		apiId = append(apiId, v.Id)
	}
	if len(api) != len(param.Apis) {
		return xerr.NewErrCodeMsg(xerr.ValidateParamErr, "apis参数错误")
	}

	oldApi, err := u.db.Sys().GetApiByMenuIds([]int{menuId})
	if err != nil {
		return err
	}
	oldApiId := make([]int, 0, len(oldApi))
	for _, v := range oldApi {
		oldApiId = append(oldApiId, v.Id)
	}

	//获取菜单信息
	ids := []int{param.ParentId}
	menu, err := u.db.Sys().GetMenuByIds(ids)
	if err != nil {
		return err
	}
	parentPath := ""
	if len(menu) > 0 {
		parentPath += menu[0].IdPath + "/" + strconv.Itoa(menu[0].MenuId)
	} else {
		parentPath = "/0"
	}

	deleteApi, addApi := u.compareApi(oldApiId, apiId)
	return u.db.Sys().UpdateMenu(c.GetInt("user_id"), "parentPath", param, deleteApi, addApi)
}

func (u *SysSrv) compareApi(oldApiId, apiId []int) ([]int, []int) {
	apiIdMap := make(map[int]struct{})
	oldApiIdMap := make(map[int]struct{})
	for _, v := range apiId {
		apiIdMap[v] = struct{}{}
	}

	for _, v := range oldApiId {
		oldApiIdMap[v] = struct{}{}
	}

	deleteApi := make([]int, 0)
	addApi := make([]int, 0)
	for _, v := range oldApiId {
		if _, ok := apiIdMap[v]; !ok {
			deleteApi = append(deleteApi, v)
		}
	}

	for _, v := range apiId {
		if _, ok := oldApiIdMap[v]; !ok {
			addApi = append(addApi, v)
		}
	}

	return deleteApi, addApi
}

// AddMenu 添加菜单
func (u *SysSrv) AddMenu(c *gin.Context, param *dto.AddMenuParam) error {
	//判断菜单重复
	m, err := u.db.Sys().GetMenuByField("menu_name", param.MenuName)
	if err != nil {
		return err
	}
	if m.MenuId > 0 {
		return xerr.NewErrCode(xerr.DataHasExist)
	}

	//判断apis的合法性
	api, err := u.db.Sys().GetApiByIds(param.Apis)
	if err != nil {
		return err
	}

	if len(api) != len(param.Apis) {
		return xerr.NewErrCodeMsg(xerr.ValidateParamErr, "apis参数错误")
	}
	//获取父级菜单
	parentPath := "/0"
	if param.ParentId > 0 {
		ids := []int{param.ParentId}
		if menu, err := u.db.Sys().GetMenuByIds(ids); err != nil {
			return errors.Wrap(err, "获取父级菜单错误")
		} else {
			for _, v := range menu {
				parentPath = v.IdPath + "/" + strconv.Itoa(v.MenuId)
			}
		}
	}

	if err = u.db.Sys().AddMenu(c.GetInt("user_id"), parentPath, param); err != nil {
		return err
	}
	return nil
}

func (u *SysSrv) MenuTree() ([]dto.MenuTree, error) {
	menuTree := make([]dto.MenuTree, 0)
	//获取所有子菜单
	menu, err := u.db.Sys().GetAllMenu()
	if err != nil {
		return nil, err
	}

	api, err := u.db.Sys().GetMenuApi()
	if err != nil {
		return nil, err
	}
	apiMap := make(map[int][]models.SysApi)
	for _, v := range api {
		apiMap[v.MenuId] = append(apiMap[v.MenuId], models.SysApi{
			Id:     v.MenuId,
			Handle: v.Handle,
			Title:  v.Title,
			Path:   v.Path,
			Method: v.Method,
		})
	}
	for _, v := range menu {
		menuTree = append(menuTree, dto.MenuTree{
			MenuId:   v.MenuId,
			MenuName: v.MenuName,
			Title:    v.Title,
			Icon:     v.Icon,
			Path:     v.Path,
			MenuType: v.MenuType,
			ParentId: v.ParentId,
			IdPath:   v.IdPath,
			Sort:     v.Sort,
			IsShow:   v.IsShow,
			Status:   v.Status,
			Remark:   v.Remark,
			Children: nil,
			Api:      apiMap[v.MenuId],
		})
	}
	mi := make(map[int]dto.MenuTree)
	for _, item := range menuTree {
		mi[item.MenuId] = item
	}
	resp := make([]dto.MenuTree, 0)
	for _, item := range menuTree {
		if item.ParentId == 0 {
			item.Children = u.getChildren(item.MenuId, menuTree)
			resp = append(resp, item)
			continue
		}
	}

	return resp, nil
}

func (u *SysSrv) getChildren(parentId int, menu []dto.MenuTree) []dto.MenuTree {
	resp := make([]dto.MenuTree, 0)
	for _, v := range menu {
		if v.ParentId == parentId {
			v.Children = u.getChildren(v.MenuId, menu)
			resp = append(resp, v)
		}
	}

	return resp
}
