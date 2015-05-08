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

	/*
		o := orm.NewOrm()
		o.Using("mysql")
		var lists []orm.ParamsList
		num, err := o.Raw("SELECT * FROM appUpload ").ValuesList(&lists)

		beego.Debug(num, err)
	*/
	//乱写一通了
	var Apps []*app_upload.AppUpload
	app_upload.AppLists().Filter("applisted", "1").Filter("public", "1").All(&Apps)

	for _, app := range Apps {
		app.IconUrl = app_func.GetUploadPath("icon", app.Appkey)
		app.Category = app_func.CateTran(app.Category)
		app.DownLoadUrl = app_func.GetAppDownLoadUrl(app.Appid)
		zip_url := app_func.GetUploadPath("zip", app.Appkey)
		beego.Debug(zip_url)
		app.Size = app_func.GetFileSize(zip_url)
		app.Screens = app_func.GetUploadScreensPath(app.Appkey)

	}

	beego.Debug(Apps)
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
