package main

import (
	_ "appcenter/docs"
	_ "appcenter/routers"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"
)

func main() {

	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"
		beego.SetLogger("file", `{"filename":"logs/test.log"}`)
	}
	beego.Run()
}
