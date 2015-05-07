package controllers

import (
	_ "appcenter/common/app_mysql"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Get() {
	beego.AutoRender = true

	o := orm.NewOrm()
	o.Using("mysql")
	var lists []orm.ParamsList
	num, err := o.Raw("SELECT * FROM user ").ValuesList(&lists)

	beego.Debug(num, err)

	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "index.tpl"
}
