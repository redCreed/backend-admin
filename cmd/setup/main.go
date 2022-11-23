package setup

import (
	"back-admin/api/models"
	"back-admin/internal/app/common/driver"
	"back-admin/internal/setup"
	"back-admin/pkg/db"
	"back-admin/store/impl"
	"fmt"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var (
	StartCmd   *cobra.Command
	configPath string
	logPath    string
)

func init() {
	StartCmd = &cobra.Command{
		Use:     "migrate",
		Short:   "Initialize the database",
		Example: "bkAdmin migrate -c configs/config.yaml",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}

	StartCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "configs/config.yaml", "Start server with provided configuration file")
	StartCmd.PersistentFlags().StringVarP(&logPath, "log", "l", "logs", "Specify the path for logs")

}

func run() {
	var (
		err  error
		gorm *gorm.DB
	)
	if err := driver.ReadConfig(configPath); err != nil {
		panic(err)
	}
	database := driver.Conf.Database
	tables := []interface{}{&models.SysApi{}, &models.SysLoginLog{}, &models.SysOperaLog{}, &models.SysMenu{},
		&models.SysRole{}, &models.SysUser{}, &models.SysCasbinRule{}, &models.SysUserRole{}, &models.SysRoleMenu{},
		&models.SysMenuApi{}}
	//连接数据库
	if gorm, err = db.Init(database.Driver, database.Host, database.User, database.Password, database.Database, database.Port); err != nil {
		fmt.Println("Connection specified data is rejected: ", err)
		return
	}
	sqlPath := "configs/pg.sql"
	switch database.Driver {
	case "postgres":
		//判断并创建schema
		db := impl.NewFactory(gorm)
		srv := setup.NewSrv(db)
		if err := srv.Setup(tables, sqlPath); err != nil {
			fmt.Println("创建表失败:", err)
			return
		}
		fmt.Println("初始化项目数据库成功！")
	case "mysql":
		db := impl.NewFactory(gorm)
		srv := setup.NewSrv(db)
		if err := srv.Setup(tables, sqlPath); err != nil {
			fmt.Println("创建表失败:", err)
			return
		}
		fmt.Println("初始化项目数据库成功！")
	default:
		fmt.Println("database type err ,drive:", database.Driver)
		return
	}

}
