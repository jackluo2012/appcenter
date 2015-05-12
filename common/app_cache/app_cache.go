package app_cache

import (
	"errors"
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
		case int:
			key = key + strconv.Itoa(c)
		case int64: //如果是数字
			key = key + strconv.FormatInt(c, 10)
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

/*
func SHe() {
	MyRedis := app_redis.Connect()
	defer MyRedis.RedisConn.Close()
	err := MyRedis.HSet("u:1212", "hs", "asdfasdf")
	if err != nil {
		panic(err)
	}
	//MyRedis.Do("EXEC")
	fmt.Println("HSET COMPLETE")

		row, err := MyRedis.HGet("u:1212", "hs")
		if err != nil {
			panic(err)
		}
		fmt.Println("HGET COMPLETE")
		fmt.Printf("(%T) : row = %s\n", row, row)

}
func GHe() {
	MyRedis := app_redis.Connect()
	defer MyRedis.RedisConn.Close()
	row, err := MyRedis.HGet("u:1212", "hs")
	if err != nil {
		panic(err)
	}
	fmt.Println("HGET COMPLETE")
	fmt.Printf("(%T) : row = %s\n", row, row)
}
*/
/**
 * 根据类型获取 值

//func GetAppsList() (map[string]interface{}, error) {
func GetAppsList() (mmap map[string]interface{}, le int, err error) {
	MyRedis := app_redis.Connect()
	defer MyRedis.RedisConn.Close()
	le, err = MyRedis.HLen(APPCENTERLIST)
	mmap, err = MyRedis.HGetAll(APPCENTERLIST)
	return
}
func SetAppsList(field int64, value interface{}) {
	MyRedis := app_redis.Connect()
	defer MyRedis.RedisConn.Close()
	MyRedis.HSet(APPCENTERLIST, field, value)

}

//*/
/**
func SetHget() {
	beego.Debug(APPCENTERLIST)
	rConn := app_redis.Conn()
	err := rConn.Send("SET", "foo", "bar")
	beego.Debug(err)
	//rConn.Send("GET", "foo", "bar")
	//rConn.Do("SET", "app1", "test1")
}
//*/
/*
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
*/
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
