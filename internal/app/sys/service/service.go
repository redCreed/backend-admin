package service

import (
	"back-admin/api/models"
	"back-admin/internal/app/sys/dto"
	"back-admin/store"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

type SysSrvInter interface {
	UserLogin(c *gin.Context, param *dto.UserLoginReq) (*dto.UserLoginResp, error)
	Refresh(param *dto.UserToken) (*dto.UserToken, error)
	UserList(param *dto.UserListParam) (*dto.UserListResp, error)
	UpdateUserRole(param *dto.UpdateUserRoleReq, userId int) error
	AddUser(c *gin.Context, param *dto.AddUserReq) error
	UpdateUser(c *gin.Context, param *dto.UpdateUserReq, userId int) error
	DeleteUser(c *gin.Context, userId int) error

	RoleList(param *dto.RoleListParam) (*dto.RoleListResp, error)
	AddRole(c *gin.Context, param *dto.AddRoleParam) error
	UpdateRole(c *gin.Context, param *dto.UpdateRoleParam, roleId int) error
	DeleteRole(c *gin.Context, roleId int) error

	AddMenu(c *gin.Context, param *dto.AddMenuParam) error
	UpdateMenu(c *gin.Context, param *dto.UpdateMenuParam, menuId int) error
	DeleteMenu(c *gin.Context, menuId int) error
	MenuTree() ([]dto.MenuTree, error)

	ApiList() ([]models.SysApi, error)
	UpdateApi(c *gin.Context, id int, api *models.SysApi) error
	GenerateApi(data map[string]dto.Api) error
}

type SysSrv struct {
	db       store.Db
	enforcer *casbin.SyncedEnforcer
}

func NewSysSrv(db store.Db, en *casbin.SyncedEnforcer) SysSrvInter {
	return &SysSrv{
		db:       db,
		enforcer: en,
	}
}
