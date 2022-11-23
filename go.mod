module back-admin

go 1.16

require (
	github.com/casbin/casbin/v2 v2.57.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fsnotify/fsnotify v1.5.4
	github.com/gin-gonic/gin v1.8.1
	github.com/go-playground/locales v0.14.0
	github.com/go-playground/universal-translator v0.18.0
	github.com/go-playground/validator/v10 v10.10.0
	github.com/gomodule/redigo v1.8.9
	github.com/google/uuid v1.3.0
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.6.0
	github.com/spf13/viper v1.13.0
	github.com/swaggo/files v0.0.0-20220728132757-551d4a08d97a
	github.com/swaggo/gin-swagger v1.5.3
	go.uber.org/zap v1.23.0
	golang.org/x/crypto v0.0.0-20221005025214-4161e89ecf1b
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gorm.io/driver/mysql v1.3.2
	gorm.io/driver/postgres v1.4.5
	gorm.io/driver/sqlserver v1.4.1
	gorm.io/gorm v1.24.1
	gorm.io/plugin/dbresolver v1.3.0
)

require (
	github.com/alibaba/sentinel-golang v1.0.4
	github.com/alibaba/sentinel-golang/pkg/adapters/gin v0.0.0-20221011112204-0d804bbadda5
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/fvbock/endless v0.0.0-20170109170031-447134032cb6
	github.com/sony/sonyflake v1.1.0
	github.com/swaggo/swag v1.8.7
)
