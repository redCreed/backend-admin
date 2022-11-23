package app

import (
	_ "back-admin/docs"
	"back-admin/internal/app/common"
	"back-admin/internal/app/common/driver"
	sys "back-admin/internal/app/sys/server"
	"context"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Service struct {
	r *gin.Engine
}

func New(r *gin.Engine) *Service {
	s := &Service{
		r: r,
	}
	return s
}

func (s *Service) Start(ctx context.Context) error {
	//初始化swag
	s.swag()
	s.api()
	//目前v1版本
	rootGroup := s.r.Group("v1")
	//添加路由组
	s.createRouter(
		sys.NewServer(rootGroup, s.r),
	)
	//v2
	return nil
}

func (s *Service) createRouter(option ...common.RouteOption) {
	for _, r := range option {
		r()
	}
}

func (s *Service) swag() {
	//当执行命令的位置与main.go在在同一文件夹，可以使用 -g来指定main.go
	//swag init -g cmd/bkAdmin/main.go
	if !driver.Instance.IsProd {
		s.r.RouterGroup.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}

func (s *Service) api() {

}
