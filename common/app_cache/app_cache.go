package app_cache

import (
	_ "appcenter/common/app_redis"
	"errors"
	"github.com/garyburd/redigo/redis"
	"strconv"
)

const (
	Split = ":"
)

type CacheInfo struct {
	Key    string
	Subkey interface{}
}

func GetKey(ci CacheInfo) (key string, err error) {
	key = ci.Key + Split
	//断言
	if ci.Subkey != nil {
		switch c := ci.Subkey.(type) {

		case int: //如果是数字
			key = key + Split + strconv.Itoa(c) + Split
		case []string:
			var tmp string
			for _, u := range c {
				// 将接口转换成string
				tmp += u + Split
			}
			key = key + tmp
		default:
			return "", errors.New("subkey type error!!!")
		}
	}

	return key, nil
}

func HGetAll(key string) (map[string]interface{}, error) {
	rConn := app_redis.Conn()
	rowRed, errRed := rConn.Do("HGETALL", key)
	if errRed == redis.ErrNil {
		return nil, errRed
	}
	row := r.GetRedisReply(rowRed, errRed, []string{})
	return row, errRed
}

func Hset(name string, key string, val interface{}) {
	rConn := app_redis.Conn()

}

/*
func GetKey(ci CacheInfo) (key string, err error) {
	key = ci.Key + Split
	//断言
	if ci.Subkey != nil {
		switch c := ci.Subkey.(type) {
		case int: //如果是数字
			key = key + Split + strconv.Itoa(c) + Split
		case []interface{}:
			tmp := ""
			for _, u := range c {
				// 将接口转换成string
				tmp += u.(string) + Split
			}
			key = key + tmp
		default:
			return nil, errors.New("subkey type error!!!")
		}
	}

	return key, nil
}
*/
