package common

import (
	"back-admin/internal/app/common/driver"
	"back-admin/pkg/xerr"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"strings"
)

// ValidateParam 校验参数
func ValidateParam(c *gin.Context, param interface{}) error {
	//shouldBind 要使用form标签
	//GET请求ShouldBind 其他方式走body json
	var err error
	if c.Request.Method == "GET" {
		err = c.ShouldBind(param)
	} else {
		err = c.ShouldBindJSON(param)
	}
	if err != nil {
		fmt.Println("should bind err:", err)
		return xerr.NewErrCodeMsg(xerr.ValidateParamErr, "参数非法")
	}

	//获取验证器
	valid := driver.Instance.Validator.Validate
	//获取翻译器
	trans := driver.Instance.Validator.Translator
	err = valid.Struct(param)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		sliceErrs := make([]string, 0)
		for _, e := range errs {
			sliceErrs = append(sliceErrs, e.Translate(trans))
		}
		return xerr.NewErrCodeMsg(xerr.ValidateParamErr, strings.Join(sliceErrs, ","))
	}
	return nil
}
