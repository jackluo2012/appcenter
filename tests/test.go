package main

import (
	/*	"appcenter/common/app_cache"
		. "appcenter/common/app_ckey"
		"appcenter/common/app_error"
	*/"fmt"
	"path/filepath"
	//"io/ioutil"
)

func main() {
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
