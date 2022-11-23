package store

import (
	"back-admin/api/models"
	"back-admin/internal/app/sys/dto"
)

type Sys interface {
	AddUser(u *models.SysUser) error
	UpdateUser(u *models.SysUser) error
	DeleteUser(userId int) error
	GetUserById(userId int) (*models.SysUser, error)
	GetUserByField(field, name string) (*models.SysUser, error)
	UserList(param *dto.UserListParam) (*dto.UserListResp, error)

	GetRoleById(id int) (*models.SysRole, error)
	GetRoleByNameKey(name []string) ([]models.SysRole, error)
	GetRoleList(param *dto.RoleListParam) (*dto.RoleListResp, error)
	GetRoleIdByUserId(userId int) ([]int, error)
	DeleteRoleByUserId(userId int) error
	GetRoleByRoleId(roleId []int) ([]models.SysRole, error)
	AddUserRole(userId int, roleId []int) error
	GetRoleByField(field, key string) (*models.SysRole, error)
	AddRole(adminId int, param *dto.AddRoleParam) error
	UpdateRole(adminId int, param *dto.UpdateRoleParam, oldMenu []int) error
	GetMenuByRoleId(id int) ([]models.SysRoleMenu, error)
	DeleteRole(roleId int) error
	GetRoleMenuByMenuId(id int) ([]models.SysRoleMenu, error)

	GetMenuByIds(id []int) ([]models.SysMenu, error)
	GetMenuByField(field, name string) (*models.SysMenu, error)
	AddMenu(admin int, parentPath string, param *dto.AddMenuParam) error
	UpdateMenu(admin int, parentPath string, param *dto.UpdateMenuParam, deleteApi, addApi []int) error
	DeleteMenu(menuId int) error
	GetAllMenu() ([]models.SysMenu, error)

	GetApiByMenuIds(id []int) ([]models.SysApi, error)
	GetApiByIds(id []int) ([]models.SysApi, error)
	GetMenuApi() ([]dto.MenuApiList, error)

	ApiList() ([]models.SysApi, error)
	UpdateApi(api *models.SysApi) error
	AddApi(api []models.SysApi) error
	DeleteAllApi() error
}
