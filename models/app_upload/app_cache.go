package app_upload

import (
	"appcenter/common/app_cache"
	. "appcenter/common/app_ckey"
	"appcenter/common/app_redis"
	"github.com/garyburd/redigo/redis"

	//	"strconv"
)

/**
 *	获取app数据
 */
func GetAppsList() (apps []*AppUpload) {
	rConn := app_redis.Conn()
	defer rConn.Close()
	if keys, err := redis.Ints(rConn.Do("ZRANGE", APPCENTERKEYS, 0, -1)); err == nil {
		for _, v := range keys {
			info := app_cache.CacheInfo{APPCENTERLIST, v}
			key, _ := app_cache.GetKey(info)

			v, err := redis.Values(rConn.Do("HGETALL", key))
			if err != nil {
				panic(err)
			}

			var app AppUpload
			if err := redis.ScanStruct(v, &app); err != nil {
				panic(err)
			}
			apps = append(apps, &app)
		}
	} else {
		panic(err)
	}
	return
}

/**
 *	设置app数据
 */
func SetAppsList(field int64, aud *AppUpload) {

	rConn := app_redis.Conn()
	defer rConn.Close()
	info := app_cache.CacheInfo{APPCENTERLIST, field}
	key, _ := app_cache.GetKey(info)

	if _, err := rConn.Do("HMSET", redis.Args{}.Add(key).AddFlat(aud)...); err != nil {
		panic(err)
	}
	if index, err := rConn.Do("INCR", APPCENTERINDEX); err == nil {
		if _, err := rConn.Do("ZADD", APPCENTERKEYS, index, field); err != nil {
			panic(err)
		}
	} else {
		panic(err)
	}

}