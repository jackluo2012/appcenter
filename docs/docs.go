package docs

import (
	"encoding/json"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/swagger"
)

const (
    Rootinfo string = `{"apiVersion":"1.0.0","swaggerVersion":"1.2","apis":[{"path":"/user","description":"用户数据上报接口\n"}],"info":{"title":"应用中心接口 Test API","description":"应用中心接口暂时不需要验证","contact":"net.webjoy@gmail.com","termsOfServiceUrl":"http://www.webjoy.net"}}`
    Subapi string = `{"/user":{"apiVersion":"1.0.0","swaggerVersion":"1.2","basePath":"","resourcePath":"/user","produces":["application/json","application/xml","text/plain","text/html"],"apis":[{"path":"/app_list","description":"","operations":[{"httpMethod":"POST","nickname":"获取用户App列表","type":"","summary":"获取用户App列表","parameters":[{"paramType":"body","name":"body","description":"\"body for user get app content\"","dataType":"UserAppSearchForm","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.UserAppInfo","responseModel":"UserAppInfo"},{"code":403,"message":"body is empty","responseModel":""}]}]},{"path":"/app_add","description":"","operations":[{"httpMethod":"POST","nickname":"用户添加App接口","type":"","summary":"用户添加App接口","parameters":[{"paramType":"body","name":"body","description":"\"用户添加app数据\"","dataType":"UserAppForm","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"common.app_error.CodeInfo","responseModel":"CodeInfo"},{"code":403,"message":"body is empty","responseModel":""}]}]},{"path":"/app_update","description":"","operations":[{"httpMethod":"POST","nickname":"用户更新App接口","type":"","summary":"用户更新App接口","parameters":[{"paramType":"body","name":"body","description":"\"用户更新app数据\"","dataType":"UserAppForm","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"common.app_error.CodeInfo","responseModel":"CodeInfo"},{"code":403,"message":"body is empty","responseModel":""}]}]},{"path":"/app_del","description":"","operations":[{"httpMethod":"POST","nickname":"用户删除App接口","type":"","summary":"用户删除App接口","parameters":[{"paramType":"body","name":"body","description":"\"用户删除App接口\"","dataType":"UserAppForm","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"common.app_error.CodeInfo","responseModel":"CodeInfo"},{"code":403,"message":"body is empty","responseModel":""}]}]}],"models":{"CodeInfo":{"id":"CodeInfo","properties":{"Code":{"type":"int","description":"","format":""},"Info":{"type":"string","description":"","format":""}}},"UserAppInfo":{"id":"UserAppInfo","properties":{"appid":{"type":"int","description":"","format":""},"create_date":{"type":"\u0026{time Time}","description":"","format":""},"udid":{"type":"string","description":"","format":""},"user_id":{"type":"int","description":"","format":""},"version":{"type":"string","description":"","format":""}}}}}}`
    BasePath string= "/v1"
)

var rootapi swagger.ResourceListing
var apilist map[string]*swagger.ApiDeclaration

func init() {
	err := json.Unmarshal([]byte(Rootinfo), &rootapi)
	if err != nil {
		beego.Error(err)
	}
	err = json.Unmarshal([]byte(Subapi), &apilist)
	if err != nil {
		beego.Error(err)
	}
	beego.GlobalDocApi["Root"] = rootapi
	for k, v := range apilist {
		for i, a := range v.Apis {
			a.Path = urlReplace(k + a.Path)
			v.Apis[i] = a
		}
		v.BasePath = BasePath
		beego.GlobalDocApi[strings.Trim(k, "/")] = v
	}
}


func urlReplace(src string) string {
	pt := strings.Split(src, "/")
	for i, p := range pt {
		if len(p) > 0 {
			if p[0] == ':' {
				pt[i] = "{" + p[1:] + "}"
			} else if p[0] == '?' && p[1] == ':' {
				pt[i] = "{" + p[2:] + "}"
			}
		}
	}
	return strings.Join(pt, "/")
}
