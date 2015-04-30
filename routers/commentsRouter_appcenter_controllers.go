package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	
	beego.GlobalControllerRouter["appcenter/controllers:UserAppController"] = append(beego.GlobalControllerRouter["appcenter/controllers:UserAppController"],
		beego.ControllerComments{
			"AppList",
			`/app_list`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["appcenter/controllers:UserAppController"] = append(beego.GlobalControllerRouter["appcenter/controllers:UserAppController"],
		beego.ControllerComments{
			"AppInsert",
			`/app_add`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["appcenter/controllers:UserAppController"] = append(beego.GlobalControllerRouter["appcenter/controllers:UserAppController"],
		beego.ControllerComments{
			"AppUpdate",
			`/app_update`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["appcenter/controllers:UserAppController"] = append(beego.GlobalControllerRouter["appcenter/controllers:UserAppController"],
		beego.ControllerComments{
			"AppRemove",
			`/app_del`,
			[]string{"post"},
			nil})

}
