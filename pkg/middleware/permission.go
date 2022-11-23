package middleware

import (
	"back-admin/internal/app/common/driver"
	"back-admin/internal/app/common/response"
	"back-admin/pkg/xerr"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Permission() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//用户角色
		roles := ctx.GetStringSlice("roles")
		for _, roleName := range roles {
			if roleName == "admin" {
				ctx.Next()
				return
			}
		}

		var (
			ret bool
			err error
		)
		for _, roleName := range roles {
			ret, err = driver.Instance.Enforcer.Enforce(roleName, ctx.Request.URL.Path, ctx.Request.Method)
			if err != nil {
				driver.Instance.Log.Error("Enforce err",
					zap.String("role_name", roleName),
					zap.String("path", ctx.Request.URL.Path),
					zap.String("method", ctx.Request.Method),
					zap.Error(err))
				response.Data(ctx, nil, xerr.NewErrCode(xerr.NoPermission))
				ctx.Abort()
				return
			}

			if ret {
				break
			}

		}

		if ret {
			ctx.Next()
		} else {
			response.Data(ctx, nil, xerr.NewErrCode(xerr.NoPermission))
			ctx.Abort()
			return
		}
	}
}
