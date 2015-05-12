package controllers

import (
	"appcenter/common/app_func"
	"appcenter/models/app_upload"
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Get() {
	beego.AutoRender = true

	uid := "100001"                    //c.GetString("uid")   //字符串
	udid := "fe80::713d:36ec:e1ed:f4b" //c.GetString("udid") //用户的帐号
	secretkey := c.GetString("secretkey")
	app_func.CheckSecurity(uid, udid)

	beego.Debug(udid)
	beego.Debug(app_func.CheckSecurity(uid, secretkey))

	var Apps []*app_upload.AppUpload
	Apps = app_upload.SearchAppLists(1, uid, udid, 889)

	//beego.Debug(Apps)
	c.Data["Apps"] = Apps

	str := app_func.Md5([]byte("hehehe"))
	beego.Debug(str)

	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "index.tpl"

}
func (c *IndexController) Detail() {
	appid := c.GetString(":id")
	beego.Debug(appid)
	beego.AutoRender = true
	c.TplNames = "show.tpl"
}
