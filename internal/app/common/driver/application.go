package driver

import (
	mycasbin "back-admin/pkg/casbin"
	"back-admin/pkg/db"
	"back-admin/pkg/log"
	"back-admin/pkg/queue"
	"back-admin/pkg/redis"
	"back-admin/pkg/validate"
	"context"
	"fmt"
	"github.com/casbin/casbin/v2"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Application struct {
	Log       *zap.Logger
	Cache     *redis.Pool
	Orm       *gorm.DB
	Validator Validate
	Enforcer  *casbin.SyncedEnforcer
	IsProd    bool
	Queue     queue.Queue
}

type Validate struct {
	Translator ut.Translator
	Validate   *validator.Validate
}

var (
	Instance *Application
)

//Init 初始化操作 日志、数据库、redis等
func Init(isProd bool) {
	var (
		err  error
		gorm *gorm.DB
	)
	Instance = new(Application)
	Instance.IsProd = isProd
	//init log
	zap := log.NewLog("", isProd)
	zap.Start(context.Background())
	Instance.Log = zap.Logger

	//init db
	database := Conf.Database
	if gorm, err = db.Init(database.Driver, database.Host, database.User, database.Password, database.Database, database.Port); err != nil {
		panic(err)
	}
	Instance.Orm = gorm

	//init redis
	redisUrl := fmt.Sprintf("%s:%s", Conf.Redis.Host, Conf.Redis.Host)
	Instance.Cache = redis.Init(redisUrl, 0)

	//init validator
	Instance.Validator.Translator, Instance.Validator.Validate, err = validate.New("zh")
	if err != nil {
		panic(err)
	}

	//init rbac
	enforcer, err := mycasbin.New(gorm, isProd)
	if err != nil {
		panic(err)
	}

	Instance.Enforcer = enforcer

}
