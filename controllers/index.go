package controllers

import (
	"appcenter/common/app_func"
	"appcenter/models"
	"appcenter/models/app_upload"
	"github.com/astaxie/beego"
	"strings"
)

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Get() {
	beego.AutoRender = true

	uid := c.GetString("uid")   //字符串
	udid := c.GetString("udid") //用户的帐号
	secretkey := c.GetString("secretkey")
	//app_func.CheckSecurity(uid, udid)

	//	beego.Debug(udid)
	beego.Debug(app_func.CheckSecurity(uid, secretkey))
	checked := app_func.CheckSecurity(uid, secretkey)
	var Apps []*app_upload.AppUpload
	Apps = app_upload.SearchAppLists(1)
	for _, app := range Apps {
		uai := models.GetUserAppsByUdid(uid, udid, app.Appid)
		app.Install = "1"
		if uai != nil && checked {
			if app.Version == uai.Version {
				app.Install = "2"
			} else {
				app.Install = "3"
			}
		}
	}
	//获取用户的安装信息

	//beego.Debug(Apps)
	c.Data["Apps"] = Apps
	c.Data["uid"] = uid
	c.Data["udid"] = udid
	c.Data["secretkey"] = secretkey
	c.TplNames = "index.tpl"

}
func (c *IndexController) Detail() {
	beego.AutoRender = true
	appid, err := c.GetInt64(":id")
	if err != nil {

	}
	uid := c.GetString("uid")   //字符串
	udid := c.GetString("udid") //用户的帐号
	secretkey := c.GetString("secretkey")
	beego.Debug(app_func.CheckSecurity(uid, secretkey))
	checked := app_func.CheckSecurity(uid, secretkey)

	//获取用户缓存中的数据
	uai := models.GetUserAppsByUdid(uid, udid, appid)
	//app 信息
	AppInfo := app_upload.GetAppInfoById(appid)
	//	beego.Debug("===================", AppInfo)
	if AppInfo == nil {

	} else {
		//是否安装
		AppInfo.Install = "1"
		if uai != nil && checked {
			if AppInfo.Version == uai.Version {
				AppInfo.Install = "2"
			} else {
				AppInfo.Install = "3"
			}
		}
		// 处理 图片
		Screens := strings.Split(AppInfo.Screens, ",")
		/*
			beego.Debug(appid)
			beego.Debug(udid)
			beego.Debug(checked)
		*/

		c.Data["AppInfo"] = AppInfo
		c.Data["Screens"] = Screens
	}

	c.TplNames = "show.tpl"
}
