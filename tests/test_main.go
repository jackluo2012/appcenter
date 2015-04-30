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
}
