package middleware

import (
	"back-admin/api/models"
	"back-admin/internal/app/common"
	"back-admin/internal/app/common/driver"
	"back-admin/pkg/queue"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strings"
	"time"
)

func Log() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		if c.Request.Method == http.MethodOptions {
			return
		}
		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由
		reqUri := c.Request.RequestURI
		// 状态码
		statusCode := c.Writer.Status()
		// 请求IP
		clientIP := c.ClientIP()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		title := ""
		s := strings.Split(c.Request.RequestURI, "/")
		if len(s) >= 3 {
			title = s[2]
		}

		log := &models.SysOperaLog{
			Title:         title,
			BusinessType:  "",
			BusinessTypes: "",
			Method:        c.HandlerName(),
			RequestMethod: reqMethod,
			OperatorType:  "",
			OperName:      "",
			DeptName:      "",
			OperUrl:       reqUri,
			OperIp:        clientIP,
			OperLocation:  "",
			OperParam:     "",
			Status:        0,
			OperTime:      time.Now(),
			JsonResult:    "",
			Remark:        "",
			LatencyTime:   latencyTime.String(),
			UserAgent:     c.Request.UserAgent(),
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		}

		if c.Request.Method != "OPTIONS" && statusCode != 404 {
			if d, err := json.Marshal(log); err != nil {
				fmt.Println("log marshal:", err)
				return
			} else {
				msg := &queue.Message{}
				msg.SetId(uuid.New().String())
				msg.SetKey(common.OperateLog)
				msg.SetErrorCount(0)
				msg.SetValues(d)
				//推送内存队列
				driver.Instance.Queue.Add(msg)
			}
		}
	}
}
