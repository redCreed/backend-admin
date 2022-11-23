package impl

import (
	"back-admin/api/models"
	"back-admin/internal/app/sys/dto"
	"back-admin/pkg/utils"
	"back-admin/pkg/xerr"
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

type sys struct {
	gorm *gorm.DB
}

func (d *sys) GetMenuApi() ([]dto.MenuApiList, error) {
	resp := make([]dto.MenuApiList, 0)
	err := d.gorm.Table("sys_menu_api").Select("sys_menu_api.menu_id,sys_menu_api.api_id," +
		"sys_api.path,sys_api.method,sys_api.handle,sys_api.title").
		Joins("left join sys_api on sys_menu_api.api_id=sys_api.id").Find(&resp).Error

	return resp, errors.WithStack(err)
}

func (d *sys) DeleteAllApi() error {
	if err := d.gorm.Where("1 = 1").Delete(&models.SysApi{}).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (d *sys) AddApi(api []models.SysApi) error {
	if err := d.gorm.Create(&api).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (d *sys) ApiList() ([]models.SysApi, error) {
	resp := make([]models.SysApi, 0)
	if err := d.gorm.Find(&resp).Error; err != nil {
		return resp, errors.WithStack(err)
	}

	return resp, nil
}

func (d *sys) UpdateApi(api *models.SysApi) error {
	if err := d.gorm.Where("id=?", api.Id).Updates(api).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (d *sys) DeleteRole(roleId int) error {
	return d.gorm.Transaction(func(tx *gorm.DB) error {
		var err error
		//删除role
		if err = tx.Model(&models.SysRole{}).Where("role_id=?", roleId).Update("status", -1).Error; err != nil {
			return errors.WithStack(err)
		}

		//删除user_role
		if err = tx.Where("role_id=?", roleId).Delete(&models.SysUserRole{}).Error; err != nil {
			return errors.WithStack(err)
		}
		//删除role_menu
		if err = tx.Where("role_id=?", roleId).Delete(&models.SysRoleMenu{}).Error; err != nil {
			return errors.WithStack(err)
		}

		return nil
	})
}

func (d *sys) UpdateMenu(adminId int, parentPath string, param *dto.UpdateMenuParam, deleteApi, addApi []int) error {
	//更新role
	role := &models.SysMenu{
		MenuName:  param.MenuName,
		MenuType:  param.MenuType,
		Status:    param.Status,
		ControlBy: models.ControlBy{UpdateBy: adminId},
		ModelTime: models.ModelTime{UpdatedAt: time.Now()},
	}

	fields := make([]string, 0)
	fields = append(fields, "menu_name", "updateBy", "updatedAt", "status", "menu_type", "parent_id")
	if param.Sort != nil {
		role.Sort = *param.Sort
		fields = append(fields, "sort")
	}

	if param.IsShow != nil {
		role.IsShow = *param.IsShow
		fields = append(fields, "is_show")
	}

	if param.Path != nil {
		role.Path = *param.Path
		fields = append(fields, "path")
	}

	if param.Title != nil {
		role.Title = *param.Title
		fields = append(fields, "title")
	}

	if param.Icon != nil {
		role.Icon = *param.Icon
		fields = append(fields, "icon")
	}

	if param.Remark != nil {
		role.Remark = *param.Remark
		fields = append(fields, "remark")
	}

	return d.gorm.Transaction(func(tx *gorm.DB) error {
		if err := tx.Select(fields).Where("menu_id=?", param.MenuId).Updates(role).Error; err != nil {
			return errors.WithStack(err)
		}

		if len(deleteApi) > 0 {
			if err := tx.Where("menu_id=?", param.MenuId).Delete(&models.SysMenuApi{}).Error; err != nil {
				return errors.WithStack(err)
			}
		}
		if len(addApi) > 0 {
			api := make([]models.SysMenuApi, 0, len(addApi))
			for _, v := range addApi {
				api = append(api, models.SysMenuApi{
					MenuId:    param.MenuId,
					ApiId:     v,
					ModelTime: models.ModelTime{CreatedAt: time.Now()},
				})
			}

			if err := tx.Create(api).Error; err != nil {
				return errors.WithStack(err)
			}
		}

		return nil
	})
}

func (d *sys) DeleteMenu(menuId int) error {
	return d.gorm.Transaction(func(tx *gorm.DB) error {
		var err error
		//删除菜单
		if err = tx.Model(&models.SysMenu{}).Where("menu_id=?", menuId).Update("status", -1).Error; err != nil {
			return errors.WithStack(err)
		}
		//删除role_menu
		if err = tx.Where("menu_id=?", menuId).Delete(&models.SysRoleMenu{}).Error; err != nil {
			return errors.WithStack(err)
		}

		//删除menu_api
		if err = tx.Where("menu_id=?", menuId).Delete(&models.SysMenuApi{}).Error; err != nil {
			return errors.WithStack(err)
		}

		return nil
	})
}

func (d *sys) AddMenu(admin int, parentPath string, param *dto.AddMenuParam) error {
	menu := &models.SysMenu{
		MenuName: param.MenuName,
		Title:    param.Title,
		Icon:     param.Icon,
		Path:     param.Path,
		MenuType: param.MenuType,
		ParentId: param.ParentId,
		IdPath:   parentPath,
		Sort:     param.Sort,
		IsShow:   param.IsShow,
		Status:   param.Status,
		Remark:   param.Remark,
		ControlBy: models.ControlBy{
			CreateBy: admin,
			UpdateBy: 0,
		},
		ModelTime: models.ModelTime{
			CreatedAt: time.Now(),
		},
	}
	return d.gorm.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(menu).Error; err != nil {
			return errors.Wrap(err, "创建菜单错误")
		}
		if len(param.Apis) > 0 {
			api := make([]models.SysMenuApi, 0, len(param.Apis))
			for _, v := range param.Apis {
				api = append(api, models.SysMenuApi{
					MenuId:    menu.MenuId,
					ApiId:     v,
					ModelTime: models.ModelTime{CreatedAt: time.Now()},
				})
			}

			if err := tx.Create(&api).Error; err != nil {
				return errors.Wrap(err, "创建菜单动作错误")
			}
		}
		return nil
	})
}

func (d *sys) GetAllMenu() ([]models.SysMenu, error) {
	menu := make([]models.SysMenu, 0)
	if err := d.gorm.Where("status>-1").Find(&menu).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return menu, nil
}

func (d *sys) GetApiByIds(id []int) ([]models.SysApi, error) {
	api := make([]models.SysApi, 0)
	if len(id) == 0 {
		return api, nil
	}

	err := d.gorm.Where("id in ?", id).Find(&api).Error
	return api, errors.WithStack(err)
}

func (d *sys) GetMenuByField(field, name string) (*models.SysMenu, error) {
	resp := new(models.SysMenu)
	err := d.gorm.Where(fmt.Sprintf("status=1 and %s='%s' ", field, name)).First(resp).Error
	if err != gorm.ErrRecordNotFound {
		return resp, errors.WithStack(err)
	}
	return resp, nil
}

func (d *sys) GetApiByMenuIds(id []int) ([]models.SysApi, error) {
	resp := make([]models.SysApi, 0)
	if len(id) == 0 {
		return resp, nil
	}
	err := d.gorm.Table("sys_menu_api").Select("sys_api.id,sys_api.path,sys_api.method").
		Joins("left join sys_api on sys_menu_api.api_id=sys_api.id").
		Where("sys_menu_api.menu_id in ?", id).Find(&resp).Error

	return resp, errors.WithStack(err)
}

// AddRole 新增角色
func (d *sys) AddRole(adminId int, param *dto.AddRoleParam) error {
	role := &models.SysRole{
		RoleName:  param.RoleName,
		Status:    1,
		RoleKey:   param.RoleKey,
		RoleSort:  param.RoleSort,
		Flag:      param.Flag,
		Remark:    param.Remark,
		Admin:     false,
		ControlBy: models.ControlBy{CreateBy: adminId},
		ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	return d.gorm.Transaction(func(tx *gorm.DB) error {
		var err error
		if err = tx.Create(role).Error; err != nil {
			return errors.WithStack(err)
		}

		if len(param.MenuIds) > 0 {
			menu := make([]models.SysRoleMenu, 0, len(param.MenuIds))
			for _, m := range param.MenuIds {
				menu = append(menu, models.SysRoleMenu{
					RoleId:    role.RoleId,
					MenuId:    m,
					ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()},
				})
			}
			if err = tx.Create(&menu).Error; err != nil {
				return errors.WithStack(err)
			}
		}
		return nil
	})
}

func (d *sys) UserList(param *dto.UserListParam) (*dto.UserListResp, error) {
	resp := new(dto.UserListResp)
	offset, size := utils.GetPageSize(param.PageNo, param.PageSize)
	users := make([]models.SysUser, 0)
	var err error
	if err = d.gorm.Where("status=?", 1).Offset(offset).Limit(size).Find(&users).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	var count int64 = 0
	if err = d.gorm.Where("status=?", 1).Model(&models.SysUser{}).Count(&count).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	userInfos := make([]dto.SysUserList, 0)
	userId := make([]int, 0)
	for _, v := range users {
		userInfos = append(userInfos, dto.SysUserList{
			UserId:    v.UserId,
			Username:  v.Username,
			NickName:  v.NickName,
			Phone:     v.Phone,
			Avatar:    v.Avatar,
			Sex:       v.Sex,
			Email:     v.Email,
			Remark:    v.Remark,
			Status:    v.Status,
			CreateBy:  v.CreateBy,
			UpdateBy:  v.UpdateBy,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			Roles:     make([]dto.RoleInfo, 0),
		})
		userId = append(userId, v.UserId)
	}

	resp.Count = count
	resp.Data = userInfos

	//获取用户角色
	if len(userId) > 0 {
		userRole := make([]dto.RoleInfo, 0)
		if err = d.gorm.Model(&models.SysUserRole{}).Select("sys_user_role.user_id,sys_user_role.role_id,role.role_name").
			Joins("left join sys_role role on sys_user_role.role_id = role.role_id").
			Where("sys_user_role.user_id in (?)", userId).Find(&userRole).Error; err != nil {
			return nil, errors.WithStack(err)
		}
		roleMap := make(map[int][]dto.RoleInfo)
		if len(userRole) > 0 {
			for _, v := range userRole {
				roleMap[v.UserId] = append(roleMap[v.UserId], v)
			}
			for _, v := range userInfos {
				v.Roles = roleMap[v.UserId]
			}

			for k, v := range userInfos {
				v.Roles = roleMap[v.UserId]
				userInfos[k] = v
			}

			resp.Data = userInfos
		}
	}

	return resp, errors.WithStack(err)
}

func (d *sys) GetUserByField(field, name string) (*models.SysUser, error) {
	resp := new(models.SysUser)
	err := d.gorm.Where(fmt.Sprintf("status=1 and %s='%s' ", field, name)).First(resp).Error
	if err == gorm.ErrRecordNotFound {
		return resp, xerr.NewErrCode(xerr.DataNoExist)
	}
	return resp, errors.WithStack(err)
}

func (d *sys) GetRoleById(id int) (*models.SysRole, error) {
	resp := new(models.SysRole)
	err := d.gorm.Where("status=1 and role_id=?", id).First(resp).Error
	return resp, errors.WithStack(err)
}

func (d *sys) GetRoleByNameKey(name []string) ([]models.SysRole, error) {
	resp := make([]models.SysRole, 0)
	if len(name) == 0 {
		return resp, nil
	}
	if err := d.gorm.Where("status=1 and role_key in (?)", name).Find(&resp).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return resp, nil
		}
		return nil, errors.WithStack(err)
	}

	return resp, nil
}

func (d *sys) AddUser(u *models.SysUser) error {
	//判断是否名称重复
	user, err := d.GetUserByField("username", u.Username)
	_, ok := err.(*xerr.CodeError)
	if err != nil && !ok {
		return errors.WithStack(err)
	}
	if user.UserId > 0 {
		return xerr.NewErrCode(xerr.UserHasExist)
	}

	if err = d.gorm.Create(u).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (d *sys) UpdateUser(u *models.SysUser) error {
	//判断是否名称重复
	if u.Username != "" {
		user, err := d.GetUserByField("username", u.Username)
		if err != nil {
			return errors.WithStack(err)
		}
		if user.UserId == 0 {
			return xerr.NewErrCode(xerr.DataNoExist)
		}
	}

	if err := d.gorm.Where("user_id=?", u.UserId).Updates(u).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (d *sys) DeleteUser(userId int) error {
	user := new(models.SysUser)
	user.Status = -1
	user.UpdatedAt = time.Now()
	err := d.gorm.Where("status=1 and user_id=? ", userId).Updates(user).Error

	return errors.WithStack(err)
}

func (d *sys) GetUserById(userId int) (*models.SysUser, error) {
	resp := new(models.SysUser)
	err := d.gorm.Where("status=1 and user_id=? ", userId).First(resp).Error
	if err == gorm.ErrRecordNotFound {
		return resp, xerr.NewErrCode(xerr.DataNoExist)
	}
	return resp, errors.WithStack(err)
}

func (d *sys) GetRoleList(param *dto.RoleListParam) (*dto.RoleListResp, error) {
	resp := new(dto.RoleListResp)
	offset, size := utils.GetPageSize(param.PageNo, param.PageSize)
	roles := make([]models.SysRole, 0)
	var err error
	if err = d.gorm.Where("status=?", 1).Offset(offset).Limit(size).Find(&roles).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	var count int64 = 0
	if err = d.gorm.Where("status=?", 1).Model(&models.SysRole{}).Count(&count).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	resp.Count = count
	resp.Data = roles

	return resp, err
}

func (d *sys) GetRoleIdByUserId(userId int) ([]int, error) {
	userRole := make([]models.SysUserRole, 0)
	resp := make([]int, 0)
	if err := d.gorm.Where("user_id=?", userId).Find(&userRole).Error; err != nil {
		return resp, errors.WithStack(err)
	}

	for _, i := range userRole {
		resp = append(resp, i.RoleId)
	}

	return resp, nil
}

func (d *sys) GetRoleByRoleId(roleId []int) ([]models.SysRole, error) {
	resp := make([]models.SysRole, 0)
	if len(roleId) > 0 {
		if err := d.gorm.Where("role_id in ?", roleId).Find(&resp).Error; err != nil {
			return nil, errors.WithStack(err)
		}
	}

	return resp, nil
}

func (d *sys) DeleteRoleByUserId(userId int) error {
	role := new(models.SysUserRole)
	err := d.gorm.Where("user_id=?", userId).Delete(role).Error

	return errors.WithStack(err)
}

func (d *sys) AddUserRole(userId int, roleId []int) error {
	if len(roleId) == 0 {
		return nil
	}
	userRole := make([]models.SysUserRole, 0, len(roleId))
	for _, v := range roleId {
		userRole = append(userRole, models.SysUserRole{
			UserId: userId,
			RoleId: v,
			ModelTime: models.ModelTime{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		})
	}
	err := d.gorm.Create(&userRole).Error
	return errors.WithStack(err)
}

func (d *sys) GetRoleByField(field, key string) (*models.SysRole, error) {
	resp := new(models.SysRole)
	err := d.gorm.Where(fmt.Sprintf("status>-1 and %s='%s' ", field, key)).First(resp).Error
	if err == gorm.ErrRecordNotFound {
		return resp, xerr.NewErrCode(xerr.DataNoExist)
	}
	return resp, errors.WithStack(err)
}

func (d *sys) GetMenuByIds(id []int) ([]models.SysMenu, error) {
	resp := make([]models.SysMenu, 0, len(id))
	if len(id) == 0 {
		return resp, nil
	}

	err := d.gorm.Where("status>-1 and menu_id in ? ", id).Find(&resp).Error
	return resp, errors.WithStack(err)
}

func (d *sys) UpdateRole(adminId int, param *dto.UpdateRoleParam, oldMenu []int) error {
	//更新role
	role := &models.SysRole{
		RoleName:  param.RoleName,
		Admin:     false,
		ControlBy: models.ControlBy{UpdateBy: adminId},
		ModelTime: models.ModelTime{UpdatedAt: time.Now()},
	}

	fields := make([]string, 0)
	fields = append(fields, "role_name", "updateBy", "updatedAt")
	if param.RoleSort != nil {
		role.RoleSort = *param.RoleSort
		fields = append(fields, "role_sort")
	}

	if param.Flag != nil {
		role.Flag = *param.Flag
		fields = append(fields, "flag")
	}

	if param.Remark != nil {
		role.Remark = *param.Remark
		fields = append(fields, "remark")
	}

	menuIds := make([]int, 0)
	for _, v := range param.MenuIds {
		menuIds = append(menuIds, v)
	}
	return d.gorm.Transaction(func(tx *gorm.DB) error {
		if err := tx.Select(fields).Where("role_id=?", param.RoleId).Updates(role).Error; err != nil {
			return errors.WithStack(err)
		}

		//管理员
		if len(param.MenuIds) == 0 && len(oldMenu) == 0 {
			return nil
		}

		//对比菜单信息
		deleteMenu, addMenu := d.compareMenu(oldMenu, param.MenuIds)
		if len(deleteMenu) > 0 {
			if err := tx.Where("menu_id=?", deleteMenu).Delete(&models.SysRoleMenu{}).Error; err != nil {
				return errors.WithStack(err)
			}
		}
		if len(addMenu) > 0 {
			roleMenu := make([]models.SysRoleMenu, 0, len(addMenu))
			for _, v := range addMenu {
				roleMenu = append(roleMenu, models.SysRoleMenu{
					RoleId:    param.RoleId,
					MenuId:    v,
					ModelTime: models.ModelTime{CreatedAt: time.Now()},
				})
			}

			if err := tx.Create(roleMenu).Error; err != nil {
				return errors.WithStack(err)
			}
		}

		return nil
	})
}

func (d *sys) compareMenu(oldMenuId, menuId []int) ([]int, []int) {
	menuIdMap := make(map[int]struct{})
	oldMenuIdMap := make(map[int]struct{})
	for _, v := range menuId {
		menuIdMap[v] = struct{}{}
	}

	for _, v := range oldMenuId {
		oldMenuIdMap[v] = struct{}{}
	}

	deleteMenu := make([]int, 0)
	addMenu := make([]int, 0)
	for _, v := range oldMenuId {
		if _, ok := menuIdMap[v]; !ok {
			deleteMenu = append(deleteMenu, v)
		}
	}

	for _, v := range menuId {
		if _, ok := oldMenuIdMap[v]; !ok {
			addMenu = append(addMenu, v)
		}
	}

	return deleteMenu, addMenu
}

func (d *sys) GetMenuByRoleId(id int) ([]models.SysRoleMenu, error) {
	api := make([]models.SysRoleMenu, 0)
	if err := d.gorm.Where("role_id=?", id).Find(&api).Error; err != nil {
		return api, errors.WithStack(err)
	}

	return api, nil
}

func (d *sys) GetRoleMenuByMenuId(id int) ([]models.SysRoleMenu, error) {
	api := make([]models.SysRoleMenu, 0)
	if err := d.gorm.Where("menu_id=?", id).Find(&api).Error; err != nil {
		return api, errors.WithStack(err)
	}

	return api, nil
}
