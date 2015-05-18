package controllers

import (
	"appcenter/common/app_func"
	"appcenter/models"
	"appcenter/models/app_upload"
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"os"
	"strings"
)

var Langs []string

type IndexController struct {
	beego.Controller
	i18n.Locale
}

func settingLocales() {
	// load locales with locale_LANG.ini files
	langs := "zh-TW|zh-CN"
	for _, lang := range strings.Split(langs, "|") {
		lang = strings.TrimSpace(lang)
		files := []string{"conf/" + "locale_" + lang + ".ini"}
		if fh, err := os.Open(files[0]); err == nil {
			fh.Close()
		} else {
			files = nil
		}
		if err := i18n.SetMessage(lang, "conf/"+"locale_"+lang+".ini", files...); err != nil {
			panic(err)
		}
	}
	Langs = i18n.ListLangs()
}

// setLang sets site language version.
func (c *IndexController) setLang() bool {
	isNeedRedir := false
	hasCookie := false

	// get all lang names from i18n
	beego.Debug(len(Langs))
	if len(Langs) == 0 {
		settingLocales()
	}

	///*
	beego.Debug(hasCookie)
	beego.Debug(isNeedRedir)
	//*/ // 1. Check URL arguments.
	lang := c.GetString("lang")
	//beego.Debug(lang)
	//*
	// 2. Get language information from cookies.
	if len(lang) == 0 {
		lang = c.Ctx.GetCookie("lang")
		hasCookie = true
	} else {
		isNeedRedir = true
	}

	// Check again in case someone modify by purpose.
	if !i18n.IsExist(lang) {
		lang = ""
		isNeedRedir = false
		hasCookie = false
	}

	// 3. check if isLogin then use user setting
	/*
		if len(lang) == 0 && this.IsLogin {
			lang = i18n.GetLangByIndex(this.User.Lang)
		}
		//*/
	// 4. Get language information from 'Accept-Language'.
	//beego.Debug("浏览器是什么语言", c.Ctx.Input.Header("Accept-Language"))
	//*
	if len(lang) == 0 {
		al := c.Ctx.Input.Header("Accept-Language")
		if len(al) > 4 {
			al = al[:5] // Only compare first 5 letters.
			if i18n.IsExist(al) {
				lang = al
			}
		}
	}
	//*/
	// 4. DefaucurLang language is English.
	//*
	if len(lang) == 0 {
		lang = "zh-TW"
		isNeedRedir = false
	}

	// Save language information in cookies.
	//*
	if !hasCookie {
		c.setLangCookie(lang)
	}

	// Set language properties.
	c.Data["Lang"] = lang
	c.Data["Langs"] = Langs

	c.Lang = lang
	beego.Debug("坑货的值 :", isNeedRedir)
	return isNeedRedir
	//*/
}
func (c *IndexController) setLangCookie(lang string) {
	c.Ctx.SetCookie("lang", lang, 60*60*24*365, "/", nil, nil, false)
}
func (c *IndexController) Prepare() {
	if c.setLang() {

	}
}

func (c *IndexController) Get() {
	beego.AutoRender = true
	uid := c.GetString("uid")             //字符串
	udid := c.GetString("udid")           //用户的帐号
	secretkey := c.GetString("secretkey") //安全密钥
	ctype := c.GetString("c")             //选择类型
	s := c.GetString("s")                 //要搜索的字符串
	//app_func.CheckSecurity(uid, udid)

	//	beego.Debug(udid)
	beego.Debug(app_func.CheckSecurity(uid, secretkey))
	checked := app_func.CheckSecurity(uid, secretkey)
	var Apps []*app_upload.AppUpload
	Apps = app_upload.SearchAppLists(ctype, s)
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
	c.Data["ctype"] = ctype

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
