package server

import (
	"back-admin/internal/app/common"
	"back-admin/internal/app/common/response"
	"back-admin/internal/app/sys/dto"
	"github.com/gin-gonic/gin"
	"strconv"
)

// addMenu
// @Summary     新增菜单
// @Description 新增菜单
// @Tags        menu
// @Accept      json
// @Product     json
// @Param       data body     dto.AddMenuParam true "body"
// @Success     200  {object} response.Response{}
// @Router      /v1/sys/menu [post]
// @Security    Bearer
func (s *Server) addMenu(c *gin.Context) {
	param := new(dto.AddMenuParam)
	if err := common.ValidateParam(c, param); err != nil {
		response.Data(c, nil, err)
		return
	}

	err := s.sys.AddMenu(c, param)
	response.Data(c, nil, err)
}

// getRoleList
// @Summary     菜单列表
// @Description 菜单列表
// @Tags        menu
// @Accept      json
// @Product     json
// @Param       data query    dto.RoleListParam true "query"
// @Success     200  {object} response.Response{data=dto.RoleListResp}
// @Router      /v1/sys/menu/list [get]
// @Security    Bearer
func (s *Server) getMenuList(c *gin.Context) {
	param := new(dto.RoleListParam)
	if err := common.ValidateParam(c, param); err != nil {
		response.Data(c, nil, err)
		return
	}

	resp, err := s.sys.RoleList(param)
	response.Data(c, resp, err)
}

// addMenu
// @Summary     更新菜单
// @Description 更新菜单
// @Tags        menu
// @Accept      json
// @Product     json
// @Param       data body     dto.UpdateMenuParam true "body"
// @Success     200  {object} response.Response{}
// @Router      /v1/sys/menu/:id [put]
// @Security    Bearer
func (s *Server) updateMenu(c *gin.Context) {
	param := new(dto.UpdateMenuParam)
	if err := common.ValidateParam(c, param); err != nil {
		response.Data(c, nil, err)
		return
	}
	id := c.Param("id")
	menuId, err := strconv.Atoi(id)
	if err != nil {
		response.Data(c, nil, err)
		return
	}
	err = s.sys.UpdateMenu(c, param, menuId)
	response.Data(c, nil, err)
}

// addMenu
// @Summary     删除菜单
// @Description 删除菜单
// @Tags        menu
// @Accept      json
// @Product     json
// @Success     200  {object} response.Response{}
// @Router      /v1/sys/menu/:id [delete]
// @Security    Bearer
func (s *Server) deleteMenu(c *gin.Context) {
	id := c.Param("id")
	menuId, err := strconv.Atoi(id)
	if err != nil {
		response.Data(c, nil, err)
		return
	}
	err = s.sys.DeleteMenu(c, menuId)
	response.Data(c, nil, err)
}

// addMenu
// @Summary     菜单树
// @Description 删除树
// @Tags        menu
// @Accept      json
// @Product     json
// @Success     200  {object} response.Response{}
// @Router      /v1/sys/menu/tree [get]
// @Security    Bearer
func (s *Server) menuTree(c *gin.Context) {
	data, err := s.sys.MenuTree()
	response.Data(c, data, err)
}
