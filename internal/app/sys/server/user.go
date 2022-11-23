package server

import (
	"back-admin/internal/app/common"
	"back-admin/internal/app/common/response"
	"back-admin/internal/app/sys/dto"
	"back-admin/pkg/xerr"
	"github.com/gin-gonic/gin"
	"strconv"
)

// login
// @Summary     用户登录
// @Description 用户登录接口
// @Tags        token
// @Accept      json
// @Product     json
// @Param       data body     dto.UserLoginReq true "body"
// @Success     200  {object} response.Response{}
// @Router      /v1/sys/token/login [post]
// @Security    Bearer
func (s *Server) login(c *gin.Context) {
	param := new(dto.UserLoginReq)
	if err := common.ValidateParam(c, param); err != nil {
		response.Data(c, nil, err)
		return
	}
	resp, err := s.sys.UserLogin(c, param)
	response.Data(c, resp, err)
}

// refresh
// @Summary     用户token刷新
// @Description 用户token刷新
// @Tags        token
// @Accept      json
// @Product     json
// @Param       data body     dto.UserToken true "body"
// @Success     200  {object} response.Response{}
// @Router      /v1/sys/token/refresh [post]
// @Security    Bearer
func (s *Server) refresh(c *gin.Context) {
	param := new(dto.UserToken)
	if err := common.ValidateParam(c, param); err != nil {
		response.Data(c, nil, err)
		return
	}
	resp, err := s.sys.Refresh(param)
	response.Data(c, resp, err)
}

// logout
// @Summary     用户退出
// @Description 用户退出
// @Tags        token
// @Accept      json
// @Product     json
// @Success     200 {object} response.Response{}
// @Router      /v1/sys/token/logout [post]
// @Security    Bearer
func (s *Server) logout(c *gin.Context) {
	response.Data(c, nil, nil)
}

// updateUserRole
// @Summary     更新用户角色
// @Description 更新用户角色
// @Tags        user
// @Accept      json
// @Product     json
// @Param       data body     dto.UpdateUserRoleReq true "body"
// @Success     200  {object} response.Response{}
// @Router      /v1/sys/user/:user_id/role [put]
// @Security    Bearer
func (s *Server) updateUserRole(c *gin.Context) {
	param := new(dto.UpdateUserRoleReq)
	if err := common.ValidateParam(c, param); err != nil {
		response.Data(c, nil, err)
		return
	}

	idStr := c.Params.ByName("user_id")
	userId, err := strconv.Atoi(idStr)
	if err != nil {
		response.Data(c, nil, xerr.NewErrCode(xerr.ValidateParamErr))
		return
	}

	err = s.sys.UpdateUserRole(param, userId)
	response.Data(c, nil, err)
}

// addUser
// @Summary     新增用户
// @Description 新增用户
// @Tags        user
// @Accept      json
// @Product     json
// @Param       data body     dto.AddUserReq true "body"
// @Success     200  {object} response.Response{}
// @Router      /v1/sys/user [post]
// @Security    Bearer
func (s *Server) addUser(c *gin.Context) {
	param := new(dto.AddUserReq)
	if err := common.ValidateParam(c, param); err != nil {
		response.Data(c, nil, err)
		return
	}
	err := s.sys.AddUser(c, param)
	response.Data(c, nil, err)
}

// updateUser
// @Summary     更新用户
// @Description 更新用户
// @Tags        user
// @Accept      json
// @Product     json
// @Param       data body     dto.UpdateUserReq true "body"
// @Success     200  {object} response.Response{}
// @Router      /v1/sys/user/:user_id [put]
// @Security    Bearer
func (s *Server) updateUser(c *gin.Context) {
	param := new(dto.UpdateUserReq)
	if err := common.ValidateParam(c, param); err != nil {
		response.Data(c, nil, err)
		return
	}

	idStr := c.Params.ByName("user_id")
	userId, err := strconv.Atoi(idStr)
	if err != nil {
		response.Data(c, nil, xerr.NewErrCode(xerr.ValidateParamErr))
		return
	}
	err = s.sys.UpdateUser(c, param, userId)
	response.Data(c, nil, err)
}

// deleteUser
// @Summary     删除用户
// @Description 删除用户
// @Tags        user
// @Accept      json
// @Product     json
// @Success     200 {object} response.Response{}
// @Router      /v1/sys/user/:user_id [delete]
// @Security    Bearer
func (s *Server) deleteUser(c *gin.Context) {
	idStr := c.Params.ByName("user_id")
	userId, err := strconv.Atoi(idStr)
	if err != nil {
		response.Data(c, nil, xerr.NewErrCode(xerr.ValidateParamErr))
		return
	}
	err = s.sys.DeleteUser(c, userId)
	response.Data(c, nil, err)
}

// getUserList
// @Summary     用户列表
// @Description 用户列表
// @Tags        user
// @Accept      json
// @Product     json
// @Param       data query    dto.UserListParam true "query"
// @Success     200  {object} response.Response{data=dto.UserListResp}
// @Router      /v1/sys/user/list [get]
// @Security    Bearer
func (s *Server) getUserList(c *gin.Context) {
	param := new(dto.UserListParam)
	if err := common.ValidateParam(c, param); err != nil {
		response.Data(c, nil, err)
		return
	}

	resp, err := s.sys.UserList(param)
	response.Data(c, resp, err)
}
