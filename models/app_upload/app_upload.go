package app_upload

import (
	//"appcenter/common/app_cache"
	//	. "appcenter/common/app_ckey"
	"appcenter/common/app_func"
	_ "appcenter/common/app_mysql"
	"appcenter/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(AppUpload))
}

type AppUpload struct {
	Appid          int64  `orm:"pk" redis:"id"`
	Name           string `redis:"m"`
	Account        string `redis:"at"`
	Category       string `redis:"cg"`
	Appkey         string `redis:"ak"`
	Description    string `redis:"dp"`
	PraiseCounts   int64  `redis:"pc"`
	DownloadCounts int64  `redis:"dlc"`
	Version        string `redis:"v"`
	Applisted      string `redis:"ald"`
	Public         string `redis:"pc"`
	Author         string `redis:"ah"`
	Source         string `redis:"sc"`
	Level          string `redis:"lv"`
	Created        string `redis:"cd"`
	Screens        string `orm:"-" redis:"ss"`
	Size           string `orm:"-" redis:"sz"`
	IconUrl        string `orm:"-" redis:"iu"`
	Install        uint8  `orm:"-" redis:"il"`
	DownLoadUrl    string `orm:"-" redis:"dlu"`
	/*
		AndroidUrl     string `orm:"-"`
		IponeUrl       string `orm:"-"`
		ElastosUrl     string `orm:"-"`
		Zip            string `orm:"-"`
	*/
}

//var apps []AppUpload
func SearchAppLists(ct int, uid string, udid string, appid int64) (apps []*AppUpload) {
	if ct == 1 {

	}
	//var uai *models.UserAppInfo
	//获取用户的信息
	beego.Debug(models.GetUserAppsByUdid(uid, udid, appid))
	apps = GetAppsList()
	if len(apps) == 0 {
		AppLists().Filter("applisted", "1").Filter("public", "1").All(&apps)
		for _, app := range apps {
			app.IconUrl = app_func.GetUploadPath("icon", app.Appkey)
			app.Category = app_func.CateTran(app.Category)
			app.DownLoadUrl = app_func.GetAppDownLoadUrl(app.Appid)
			zip_url := app_func.GetUploadPath("zip", app.Appkey)
			app.Size = app_func.GetFileSize(zip_url)
			app.Install = 1
			app.Screens = app_func.GetUploadScreensPath(app.Appkey)
			SetAppsList(app.Appid, app)
		}
	}

	return
}

func (a *AppUpload) TableName() string {
	return "appUpload"
}

func AppLists() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(AppUpload))
}

func AppList() {
	o := orm.NewOrm()
	o.Using("mysql")
	var lists []orm.ParamsList
	num, err := o.Raw("SELECT * FROM appUpload ").ValuesList(&lists)

	beego.Debug(num, err)
}
