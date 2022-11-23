package server

import (
	"back-admin/api/models"
	"back-admin/internal/app/common"
	"back-admin/internal/app/common/response"
	"back-admin/internal/app/sys/dto"
	"github.com/gin-gonic/gin"
	"strconv"
)

// getApiList
// @Summary     接口列表
// @Description 接口列表
// @Tags        api
// @Accept      json
// @Product     json
// @Success     200  {object} response.Response{data=[]models.SysApi}
// @Router      /v1/sys/api/list [get]
// @Security    Bearer
func (s *Server) getApiList(c *gin.Context) {
	resp, err := s.sys.ApiList()
	response.Data(c, resp, err)
}

// updateApi
// @Summary     更新api接口
// @Description 更新api接口
// @Tags        api
// @Accept      json
// @Product     json
// @Param       data body    models.SysApi true "body"
// @Success     200  {object} response.Response{ }
// @Router      /v1/sys/api/:id [put]
// @Security    Bearer
func (s *Server) updateApi(c *gin.Context) {
	param := new(models.SysApi)
	if err := common.ValidateParam(c, param); err != nil {
		response.Data(c, nil, err)
		return
	}

	id := c.Param("id")
	apiId, err := strconv.Atoi(id)
	if err != nil {
		response.Data(c, nil, err)
		return
	}
	err = s.sys.UpdateApi(c, apiId, param)
	response.Data(c, nil, err)
}

// updateApi
// @Summary     生成接口
// @Description 生成接口
// @Tags        api
// @Accept      json
// @Product     json
// @Success     200  {object} response.Response{ }
// @Router      /v1/sys/api/generate [post]
// @Security    Bearer
func (s *Server) generateApi(c *gin.Context) {
	apiMap := make(map[string]dto.Api)
	routers := s.engin.Routes()
	for _, v := range routers {
		if v.Path != "/swagger/*any" {
			apiMap[v.Path+"_"+v.Method] = dto.Api{
				Handler:     v.Handler,
				Method:      v.Method,
				Description: "",
				Uri:         v.Path,
			}
		}

	}

	err := s.sys.GenerateApi(apiMap)
	response.Data(c, nil, err)
}
