package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

func Init(driver, host, user, password, database, port string) (*gorm.DB, error) {
	dsn := ""
	var dial gorm.Dialector
	postgres.Open(dsn)
	switch driver {
	case "postgres":
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
			host, user, password, database, port)
		dial = postgres.Open(dsn)
	default:
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			user, password, host, port, database)
		dial = mysql.Open(dsn)
	}

	//newLogger := logger.New(
	//	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	//	logger.Config{
	//		SlowThreshold: time.Second,   // 慢 SQL 阈值
	//		LogLevel:      logger.Silent, // Log level
	//		Colorful:      false,         // 禁用彩色打印
	//	},
	//)

	db, err := gorm.Open(dial, &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
		QueryFields: true, //避免select *
	})
	if err != nil {
		return nil, err
	}

	gdb, err := db.DB()
	if err != nil {
		return nil, err
	}
	gdb.SetConnMaxLifetime(240 * time.Second)
	gdb.SetMaxIdleConns(100)
	gdb.SetMaxOpenConns(100)

	return db, nil
}
