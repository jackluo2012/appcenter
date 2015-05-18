package app_upload

import (
	//"appcenter/common/app_cache"
	//	. "appcenter/common/app_ckey"
	"appcenter/common/app_func"
	_ "appcenter/common/app_mysql"
	//	"appcenter/models"
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
	Apptype        string
	Author         string `redis:"ah"`
	Source         string `redis:"sc"`
	Level          string `redis:"lv"`
	Created        string `redis:"cd"`
	Screens        string `orm:"-" redis:"ss"`
	Size           string `orm:"-" redis:"sz"`
	IconUrl        string `orm:"-" redis:"iu"`
	Install        string `orm:"-" redis:"il"`
	DownLoadUrl    string `orm:"-" redis:"dlu"`
	/*
		AndroidUrl     string `orm:"-"`
		IponeUrl       string `orm:"-"`
		ElastosUrl     string `orm:"-"`
		Zip            string `orm:"-"`
	*/
}

//var apps []AppUpload
// ct 是类型
// s 要搜索的字符串
func SearchAppLists(ct string, s string) (apps []*AppUpload) {
	apps = GetAppsList()
	if len(apps) == 0 {
		AppLists().Filter("applisted", "1").Filter("public", "1").All(&apps)
		for _, app := range apps {
			app.IconUrl = app_func.GetUploadPath("icon", app.Appkey)
			app.Category = app_func.CateTran(app.Category)
			app.DownLoadUrl = app_func.GetAppDownLoadUrl(app.Appid)
			zip_url := app_func.GetUploadPath("zip", app.Appkey)
			app.Size = app_func.GetFileSize(zip_url)
			app.Screens = app_func.GetUploadScreensPath(app.Appkey)
			SetAppsList(app.Appid, app)
		}
	}
	if s != "" {
		apps = FindAppByName(s)
	}
	if ct == "1" {
		return
	}
	if ct == "0" {
		apps = FindAppByType(ct)
	}

	//var uai *models.UserAppInfo
	//获取用户的信息
	//	beego.Debug(models.GetUserAppsByUdid(uid, udid, appid))
	//缓存中拿 数据

	return
}

/**
 * 暂时先这样子写了
 */
func GetAppInfoById(appid int64) (app *AppUpload) {
	app = GetAppCacheInfoById(appid)
	//beego.Debug("缓存中的数据 ", app)
	if app == nil {
		var app_info AppUpload
		AppLists().Filter("appid", appid).One(&app_info)
		//beego.Debug("取出的值", appid, "====>", app_info)
		if app_info.Appid > 0 {
			app = &app_info
		}
		//		beego.Debug("数据库中的数据", app)
	}
	return
}

func (a *AppUpload) TableName() string {
	return "appUpload"
}

func AppLists() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(AppUpload))
}

/**
 * 获取 查找列表
 */
func FindAppByName(s string) (apps []*AppUpload) {
	var maps []orm.Params
	num, err := AppLists().Filter("applisted", "1").Filter("public", "1").Filter("name__icontains", s).Values(&maps, "appid")
	//	beego.Debug("搜索结果===", num)
	//	beego.Debug(maps)
	if err == nil {
		if num > 0 {
			apps = SearchAppsList(maps)
		}
	}
	return
}

/**
 * 获取 查找列表
 */
func FindAppByType(ctype string) (apps []*AppUpload) {
	var maps []orm.Params

	num, err := AppLists().Filter("applisted", "1").Filter("public", "1").Filter("apptype", ctype).Values(&maps, "appid")
	//	beego.Debug("搜索结果===", num)
	//	beego.Debug(maps)
	if err == nil {
		if num > 0 {
			apps = SearchAppsList(maps)
		}
	}
	return
}

func AppList() {
	o := orm.NewOrm()
	o.Using("mysql")
	var lists []orm.ParamsList
	num, err := o.Raw("SELECT * FROM appUpload ").ValuesList(&lists)

	beego.Debug(num, err)
}
