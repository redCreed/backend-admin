package app

import (
	"back-admin/api/models"
	"back-admin/internal/app"
	"back-admin/internal/app/common"
	"back-admin/internal/app/common/driver"
	"back-admin/pkg/middleware"
	"back-admin/pkg/queue"
	"context"
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var (
	StartCmd = &cobra.Command{
		Use:          "run",
		Short:        "run api server",
		Example:      "bkAdmin run -c configs/config.yml",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			startHttpServer()
			return nil
		},
	}

	configPath string
	isProd     bool
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "configs/config.yaml", "Start server with provided configuration file")
	StartCmd.PersistentFlags().BoolVarP(&isProd, "mode", "m", false, "server run mode. default not prod  ")
	//解析配置
	if err := driver.ReadConfig(configPath); err != nil {
		panic(err)
	}

}

func runQueue() {
	//注册内存队列
	memory := queue.NewMemory(2000)
	driver.Instance.Queue = memory
	memory.Register(common.LoginLog, models.SaveLoginLog)
	memory.Register(common.OperateLog, models.SaveOperaLog)
}

func startHttpServer() {
	//驱动初始化
	driver.Init(isProd)
	//初始化内存队列
	runQueue()
	//初始化gin
	engin := gin.Default()
	if isProd {
		gin.SetMode(gin.ReleaseMode)
	}

	engin.Use(
		middleware.Cors(),
		middleware.Sentinel(),
		middleware.Log(),
	)

	//开启pprof监听
	go func() {
		//http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", driver.Conf.Http.Pprof.Port), nil)
		if err := endless.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", driver.Conf.Http.Pprof.Port), nil); err != nil {
			fmt.Println("pprof err:", err)
		}
	}()

	srv := app.New(engin)
	srv.Start(context.Background())

	/*
		endless平滑重启
		1、启动服务
		2、找到服务的pid，然后kill -1
			ps -ef | grep main
			kill -1 pid
		3、重启服务
	*/
	if err := endless.ListenAndServe(fmt.Sprintf(":%s", driver.Conf.Http.Port), engin); err != nil {
		panic(err)
	}
	//if err := engin.Run(fmt.Sprintf(":%s", driver.Conf.Http.Port)); err != nil {
	//	fmt.Println("start http service:", err)
	//}
	//quit := make(chan os.Signal)
	//signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	//<-quit
}
