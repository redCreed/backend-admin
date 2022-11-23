package server

import (
	"back-admin/internal/app/common"
	"back-admin/internal/app/common/response"
	"back-admin/internal/app/sys/dto"
	"github.com/gin-gonic/gin"
	"strconv"
)

// getRoleList
// @Summary     角色列表
// @Description 角色列表
// @Tags        role
// @Accept      json
// @Product     json
// @Param       data query    dto.RoleListParam true "query"
// @Success     200  {object} response.Response{data=dto.RoleListResp}
// @Router      /v1/sys/role/list [get]
// @Security    Bearer
func (s *Server) getRoleList(c *gin.Context) {
	param := new(dto.RoleListParam)
	if err := common.ValidateParam(c, param); err != nil {
		response.Data(c, nil, err)
		return
	}

	resp, err := s.sys.RoleList(param)
	response.Data(c, resp, err)
}

// addRole
// @Summary     增加角色
// @Description 增加角色
// @Tags        role
// @Accept      json
// @Product     json
// @Param       data body     dto.AddRoleParam true "body"
// @Success     200  {object} response.Response{}
// @Router      /v1/sys/role [post]
// @Security    Bearer
func (s *Server) addRole(c *gin.Context) {
	param := new(dto.AddRoleParam)
	if err := common.ValidateParam(c, param); err != nil {
		response.Data(c, nil, err)
		return
	}

	err := s.sys.AddRole(c, param)
	response.Data(c, nil, err)
}

// updateRole
// @Summary     更新角色
// @Description 更新角色
// @Tags        role
// @Accept      json
// @Product     json
// @Param       data body     dto.UpdateRoleParam true "body"
// @Success     200  {object} response.Response{}
// @Router      /v1/sys/role/:id [put]
// @Security    Bearer
func (s *Server) updateRole(c *gin.Context) {
	param := new(dto.UpdateRoleParam)
	if err := common.ValidateParam(c, param); err != nil {
		response.Data(c, nil, err)
		return
	}
	id := c.Param("id")
	roleId, err := strconv.Atoi(id)
	if err != nil {
		response.Data(c, nil, err)
		return
	}
	err = s.sys.UpdateRole(c, param, roleId)
	response.Data(c, nil, err)
}

// updateRole
// @Summary     删除角色
// @Description 删除角色
// @Tags        role
// @Accept      json
// @Product     json
// @Success     200  {object} response.Response{}
// @Router      /v1/sys/role/:id [delete]
// @Security    Bearer
func (s *Server) deleteRole(c *gin.Context) {
	id := c.Param("id")
	roleId, err := strconv.Atoi(id)
	if err != nil {
		response.Data(c, nil, err)
		return
	}
	err = s.sys.DeleteRole(c, roleId)
	response.Data(c, nil, err)
}
