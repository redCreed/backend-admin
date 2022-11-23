package server

import (
	"back-admin/internal/app/common"
	"back-admin/internal/app/common/driver"
	"back-admin/internal/app/sys/service"
	"back-admin/store/impl"
	"github.com/gin-gonic/gin"
)

/*
	系统板块
*/

type Server struct {
	engin  *gin.Engine
	router *gin.RouterGroup
	sys    service.SysSrvInter
}

func NewServer(r *gin.RouterGroup, engin *gin.Engine) common.RouteOption {
	return func() {
		db := impl.NewFactory(driver.Instance.Orm)
		srv := service.NewSysSrv(db, driver.Instance.Enforcer)
		s := &Server{engin: engin, router: r, sys: srv}
		s.initRouter()
	}
}

func (s *Server) initRouter() {
	sys := s.router.Group("/sys/token")
	{
		sys.POST("/login", s.login)
		sys.POST("/refresh", s.refresh)
		sys.POST("/logout", s.logout)
	}

	api := s.router.Group("/sys/api")
	{
		api.GET("/list", s.getApiList)
		api.PUT("/:id", s.updateApi)
		//测试环境下更新api，同时更新到线上
		if !driver.Instance.IsProd {
			api.POST("/generate", s.generateApi)
		}
	}

	userSys := s.router.Group("/sys/user")
	//.Use(middleware.Auth()) //middleware.Permission(),
	{
		//用户列表
		userSys.GET("/list", s.getUserList)
		//更新用户角色
		userSys.PUT("/:user_id/role", s.updateUserRole)
		//新增用户
		userSys.POST("", s.addUser)
		//更新用户
		userSys.PUT("/:user_id", s.updateUser)
		//删除用户
		userSys.DELETE("/:user_id", s.deleteUser)
	}

	roleSys := s.router.Group("/sys/role")
	{
		//角色列表
		roleSys.GET("/list", s.getRoleList)
		//新增角色
		roleSys.POST("", s.addRole)
		//更新角色
		roleSys.PUT("/:id", s.updateRole)
		//删除角色
		roleSys.DELETE("/:id", s.deleteRole)
	}

	menuSys := s.router.Group("/sys/menu")
	{
		//菜单列表
		menuSys.GET("/list", s.getMenuList)
		//新增菜单
		menuSys.POST("", s.addMenu)
		//更新菜单
		menuSys.PUT("/:id", s.updateMenu)
		//删除菜单
		menuSys.DELETE("/:id", s.deleteMenu)
		//菜单树
		menuSys.GET("/tree", s.menuTree)
	}

}
