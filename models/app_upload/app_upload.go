package app_upload

import (
	"appcenter/common/app_cache"
	. "appcenter/common/app_ckey"
	"appcenter/common/app_func"
	_ "appcenter/common/app_mysql"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(AppUpload))
}

type AppUpload struct {
	Appid          int64 `orm:"pk"`
	Name           string
	Account        string
	Category       string
	Appkey         string
	Description    string
	PraiseCounts   int64
	DownloadCounts int64
	Version        string
	Applisted      string
	Public         string
	Author         string
	Source         string
	Level          string
	Created        string
	Screens        []string `orm:"-"`
	Size           string   `orm:"-"`
	IconUrl        string   `orm:"-"`
	Install        int8     `orm:"-"`
	DownLoadUrl    string   `orm:"-"`
	/*
		AndroidUrl     string `orm:"-"`
		IponeUrl       string `orm:"-"`
		ElastosUrl     string `orm:"-"`
		Zip            string `orm:"-"`
	*/
}

var apps []AppUpload

func SearchAppLists(ct int) (apps []*app_upload.AppUpload) {
	if ct == 1 {

	}
	apps = AppLists().Filter("applisted", "1").Filter("public", "1").All(&Apps)

	for _, app := range Apps {
		app.IconUrl = app_func.GetUploadPath("icon", app.Appkey)
		app.Category = app_func.CateTran(app.Category)
		app.DownLoadUrl = app_func.GetAppDownLoadUrl(app.Appid)
		zip_url := app_func.GetUploadPath("zip", app.Appkey)
		beego.Debug(zip_url)
		app.Size = app_func.GetFileSize(zip_url)
		app.Screens = app_func.GetUploadScreensPath(app.Appkey)

	}
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
