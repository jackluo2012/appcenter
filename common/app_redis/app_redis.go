package app_redis

import (
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
	//	"github.com/jackluo2012/Redis"
	"time"
)

/*
func Connect() Redis.RedisType {
	//*
	server := beego.AppConfig.String("cache::server")
	password := beego.AppConfig.String("cache::password")
	port, _ := beego.AppConfig.Int("cache::port")
	db, _ := beego.AppConfig.Int("cache::db")

	red := Redis.RedisType{Host: server, Port: port, Password: password, DB: db}
	red.Connect()

	return red
}

//*/
//**
var pool *redis.Pool

func Conn() redis.Conn {
	return pool.Get()
}

func newPool(server, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if password != "" {
				_, err := c.Do("AUTH", password)
				if err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func init() {
	server := beego.AppConfig.String("cache::server")
	password := beego.AppConfig.String("cache::password")

	pool = newPool(server, password)
}

//*/
