package models

type UserAppForm struct {
	UserId  string `form:"userid"    valid:"Required"`
	AppId   int64  `form:"appid"     valid:"Required"`
	Udid    string `form:"udid" valid:"Required"`
	Version string `form:"version" `
}
type UserAppFormSlice struct {
	UserAppForms []UserAppForm `json:"lists"`
}
type UserAppSearchForm struct {
	UserId   string `form:"userid"    valid:"Required"`
	PageNo   int    `form:"pageno"     valid:"Required"`
	PageSize int    `form:"pagesize" valid:"Required"`
	AppId    int64  `form:"appid" valid:""`
	Udid     string `form:"udid" valid:""`
}

type AppInfo struct {
	Appid   int64  `form:"appid"`
	AppName string `form:"app_name"`
	Version string `form:"version"`
	Author  string `form:"author"`
	Icon    interface{}
}
