package middleware

import (
	"fmt"
	"github.com/alibaba/sentinel-golang/core/system"
	sentinel "github.com/alibaba/sentinel-golang/pkg/adapters/gin"
	"github.com/gin-gonic/gin"
)

// Sentinel 限流
func Sentinel() gin.HandlerFunc {

	if _, err := system.LoadRules([]*system.Rule{
		{
			MetricType:   system.InboundQPS,
			TriggerCount: 200,
			Strategy:     system.BBR,
		},
	}); err != nil {
		fmt.Println(err)
	}
	return sentinel.SentinelMiddleware(
		sentinel.WithBlockFallback(func(ctx *gin.Context) {
			ctx.AbortWithStatusJSON(200, map[string]interface{}{
				"msg":  "too many request; the quota used up!",
				"code": 500,
			})
		}),
	)

	return nil
}
