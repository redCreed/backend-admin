package middleware

import (
	"back-admin/internal/app/common/response"
	"back-admin/pkg/utils"
	"back-admin/pkg/xerr"
	"github.com/gin-gonic/gin"
	"time"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.Request.Header.Get("Authorization")
		if len(token) < 10 {
			response.Data(context, nil, xerr.NewErrCode(xerr.TokenMalformed))
			context.Abort()
			return
		}

		claims, err := utils.ParseToken(token)
		if err != nil {
			response.Data(context, nil, xerr.NewErrCode(xerr.TokenInvalid))
			context.Abort()
			return
		}

		if claims.ExpiresAt < time.Now().Unix() {
			response.Data(context, nil, xerr.NewErrCode(xerr.TokenNotValidYet))
			context.Abort()
			return
		}

		// 设置返回的头部
		context.Set("user_id", claims.UserId)
		context.Set("nickname", claims.Nickname)
		context.Set("roles", claims.Roles)
		context.Next()
	}
}
