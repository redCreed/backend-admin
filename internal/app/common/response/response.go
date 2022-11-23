package response

import (
	"back-admin/internal/app/common/driver"
	"back-admin/pkg/xerr"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// Response ret 0 正常  msg   data
type Response struct {
	Ret int32  `json:"ret"`
	Msg string `json:"msg"`
	//Detail string `json:"detail"`
	Data interface{} `json:"data"`
}

func Data(c *gin.Context, data interface{}, err error) {
	if data == "" || data == nil || err != nil {
		data = make([]int, 0)
	}
	resp := &Response{Data: data, Ret: xerr.OK, Msg: "success"}
	if err != nil {
		//判断是否自定义错误
		if temp, ok := errors.Cause(err).(*xerr.CodeError); ok {
			if temp.GetCode() == xerr.ValidateParamErr {
				log := fmt.Sprintf("param invalid err: %s", temp.GetMsg())
				//driver.Instance.Log.Error(log)
				fmt.Println(log)
			}
			//区分特别状态码
			resp = &Response{Ret: temp.GetCode(), Msg: temp.GetMsg(), Data: data}
		} else {
			//返回给前端
			resp = &Response{Ret: xerr.DbErr, Msg: xerr.ErrMapMsg(xerr.DbErr), Data: data}
		}

		stack := fmt.Sprintf("stack error trace:\n%+v\n", err) //错误的堆栈
		fmt.Println("err:", stack)
		driver.Instance.Log.Error(stack)
	}
	c.JSON(200, resp)
	c.Abort()
	return
}
