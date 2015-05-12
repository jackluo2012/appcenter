package main

import (
	/*	"appcenter/common/app_cache"
		. "appcenter/common/app_ckey"
		"appcenter/common/app_error"
	*/"fmt"
	"github.com/jackluo2012/Redis"
	"path/filepath"
	//"io/ioutil"
)

func connect() Redis.RedisType {
	red := Redis.RedisType{Host: "127.0.0.1", Port: 6379, Password: "", DB: 10}
	red.Connect()
	return red
}

func main() {
	MyRedis := connect()
	defer MyRedis.RedisConn.Close()
	//	sadd(MyRedis)
	//	srem(MyRedis)
	err := MyRedis.HSet("u:1212", "hs", "asdfasdf")
	if err != nil {
		panic(err)
	}
	fmt.Println("HSET COMPLETE")
	row, err := MyRedis.HGet("u:1212", "hs")
	if err != nil {
		panic(err)
	}
	fmt.Println("HGET COMPLETE")
	fmt.Printf("(%T) : row = %s\n", row, row)
	//var a interface{}
	//fmt.Println(app_error.ErrInputData)

	/*
		cacheinfo := app_cache.CacheInfo{
			Key:    UserAppList,
			Subkey: 1000,
		}
	//*/
	/*
		cacheinfo := app_cache.CacheInfo{
			Key:    UserAppList,
			Subkey: []string{"jack", "luo"},
		}
		//*/
	/*
		arr := [2]string{"jack", "luo"}
		//fmt.Println(arr)

		for _, v := range arr {
			fmt.Println(v)
		}*/
	//info := app_cache.GetKey(cacheinfo)
	//fmt.Println(info)
	//fmt.Println(a)

	/*
		fmt.Println("1212")
		files, _ := ioutil.ReadDir(".")
		for _, fi := range files {

			if fi.IsDir() {
				//listAll(path + "/" + fi.Name())
				println("." + "/" + fi.Name())
			} else {
				println("." + "/" + fi.Name())
			}
		}*/
	getUrl("*.a")
}

func getUrl(file string) {
	files, err := filepath.Glob(file)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(len(files))
}
