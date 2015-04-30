// @APIVersion 1.0.0
// @Title 应用中心接口 Test API
// @Description 应用中心接口暂时不需要验证
// @Contact net.webjoy@gmail.com
// @TermsOfServiceUrl http://www.webjoy.net
package routers

import (
	"appcenter/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserAppController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
