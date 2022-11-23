package service

import (
	"back-admin/api/models"
	"back-admin/internal/app/sys/dto"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func (u *SysSrv) ApiList() ([]models.SysApi, error) {
	return u.db.Sys().ApiList()
}

func (u *SysSrv) UpdateApi(c *gin.Context, id int, api *models.SysApi) error {
	adminId := c.GetInt("user_id")
	api.UpdatedAt = time.Now()
	api.UpdateBy = adminId
	api.Id = id
	return u.db.Sys().UpdateApi(api)
}

func (u *SysSrv) GenerateApi(apiMap map[string]dto.Api) error {
	file, err := os.Open("docs/swagger.json")
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	swag := new(dto.Swag)
	err = json.Unmarshal(data, swag)
	//k1:路径  k2:方式
	for k, v := range swag.Paths {
		value := v.(map[string]interface{})
		for k1, v1 := range value {
			value2 := v1.(map[string]interface{})
			for k2, v3 := range value2 {
				if k2 == "description" {
					uri := k + "_" + strings.ToUpper(k1)
					tempApi := apiMap[uri]
					tempApi.Description = v3.(string)
					apiMap[uri] = tempApi
				}
			}
		}
	}

	model := make([]models.SysApi, 0)
	for k, v := range apiMap {
		if k != "/v1/sys/api/generate_POST" {
			model = append(model, models.SysApi{
				Handle:    v.Handler,
				Title:     v.Description,
				Path:      v.Uri,
				Method:    v.Method,
				ModelTime: models.ModelTime{CreatedAt: time.Now()},
				ControlBy: models.ControlBy{},
			})
		}

	}
	if err := u.db.Sys().DeleteAllApi(); err != nil {
		return err
	}
	return u.db.Sys().AddApi(model)
}
