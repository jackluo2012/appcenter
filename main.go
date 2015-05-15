package main

import (
	_ "appcenter/docs"
	_ "appcenter/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	_ "github.com/astaxie/beego/session/redis"
)

func main() {

	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"
		beego.SetLogger("file", `{"filename":"logs/test.log"}`)
	}
	//ctx.Output.Header("Access-Control-Allow-Origin", "*")
	beego.InsertFilter("*", beego.BeforeRouter, func(ctx *context.Context) {
		ctx.Output.Header("Access-Control-Allow-Origin", "*")
	})
	//
	beego.InsertFilter("/v1/user/*", beego.BeforeRouter, func(ctx *context.Context) {
		//ctx.Output.Header("Access-Control-Allow-Origin", "*")
		beego.Debug("")
	})

	beego.Run()
}
