package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

type Pool struct {
	pool *redis.Pool
}

func Init(host string, library int) *Pool {
	pool := &redis.Pool{
		MaxIdle:     10, // 最大空闲连接数
		MaxActive:   10,
		IdleTimeout: 240,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp",
				host,
				redis.DialDatabase(library),
				redis.DialReadTimeout(time.Duration(1)*time.Second),
				redis.DialWriteTimeout(time.Duration(1)*time.Second),
				redis.DialConnectTimeout(time.Duration(2)*time.Second),
			)
			if err != nil {
				return nil, err
			}
			//不需要校验密码
			//if password != "" {
			//	if _, err := c.Do("AUTH", password); err != nil {
			//		c.Close()
			//		return nil, err
			//	}
			//}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if _, err := c.Do("PING"); err != nil {
				fmt.Println("PING", err)
				return err
			}
			return nil
		},
	}

	return &Pool{pool: pool}
}

func (r *Pool) Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	conn := r.pool.Get()
	reply, err = conn.Do(commandName, args...)
	defer conn.Close()
	return
}
