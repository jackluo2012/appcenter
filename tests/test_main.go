package main

import (
	"fmt"
)

type UserAppSearchForm struct {
	UserId   int    `form:"userid"    valid:"Required"`
	PageNo   int    `form:"pageno"     valid:"Required"`
	PageSize int    `form:"pagesize" valid:"Required"`
	AppId    int    `form:"appid" valid:""`
	Udid     string `form:"udid" valid:""`
}

func main() {
	u := UserAppSearchForm{}
	fmt.Println(u)
	m := map[string]map[string]string{}
	mm, ok := m["kkk"]
	if !ok {
		mm = make(map[string]string)
		m["kkk"] = mm
	}
	mm["k1k1k1"] = "sssss"
	fmt.Println(m)
}
