package controllers

import (
	"appcenter/common/app_error"
	"appcenter/models"
	"appcenter/models/app_upload"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"strconv"
	"time"
)

// 用户数据上报接口
type UserAppController struct {
	beego.Controller
}

// @Title 获取用户App列表
// @Description 获取用户App列表
// @Param	body		body 	models.UserAppSearchForm true		"body for user get app content"
// @Success 200 {object} models.UserAppInfo
// @Failure 403 body is empty
// @router /app_list [post]
func (u *UserAppController) AppList() {

	user_app := models.UserAppSearchForm{}
	json.Unmarshal(u.Ctx.Input.RequestBody, &user_app)

	beego.Debug("ParseUserAppSearchForm:", &user_app)
	valid := validation.Validation{}
	ok, err := valid.Valid(&user_app)
	//如果存在错误
	if err != nil {
		beego.Debug("ValidUserAppSearchForm:", err)
		u.Data["json"] = app_error.ErrInputData
		u.ServeJson()
		return
	}
	if !ok {
		beego.Debug("ValidUserAppSearchForm errors:")
		for _, err := range valid.Errors {
			beego.Debug(err.Key, err.Message)
		}
		u.Data["json"] = app_error.ErrDupUser
		u.ServeJson()
		return
	}

	//*
	code, lists := models.GetAllUserApps(&user_app)
	if code == -1 {
		beego.Debug("InsertUserApp:", err)
		u.Data["json"] = app_error.ErrDatabase
		u.ServeJson()
		return
	}

	for k, v := range lists {
		beego.Debug("- - -", k, v.UserId)
	}

	u.Data["json"] = lists
	u.ServeJson()
	//*/
}

// @Title 用户添加App接口
// @Description 用户添加App接口
// @Param	body		body 	models.UserAppForm	true	"用户添加app数据"
// @Success 200 {object} common.app_error.CodeInfo
// @Failure 403 body is empty
// @router /app_add [post]
func (u *UserAppController) AppInsert() {
	user_app_form := models.UserAppFormSlice{}
	json.Unmarshal(u.Ctx.Input.RequestBody, &user_app_form)
	//*
	ids := map[int64]string{}
	//uafs := []models.UserAppForm{}
	err_info := app_error.CodeInfo{}
	// 检查是否重复
	for _, uaf := range user_app_form.UserAppForms {
		//uaf.(UserAppForm)
		//beego.Debug(uaf)
		//*
		if ids[uaf.AppId] == "1" {
			err_info = app_error.ErrDataDuplication
			//return
			break
		}
		//beego.Debug(uaf.AppId)

		ids[uaf.AppId] = "1"
		//获取应用的当前版本信息

		aud := app_upload.GetAppInfoById(uaf.AppId)
		//	beego.Debug("每次都有信息...", aud)
		if aud == nil {
			err_info = app_error.ErrUserAppInfoNoExist
			//return
			break
		}
		valid := validation.Validation{}
		ok, err := valid.Valid(&uaf)
		//如果存在错误
		if err != nil {
			err_info = app_error.ErrInputData
			break
		}
		if !ok {
			err_info = app_error.ErrDupUser
			break
		}
		uaf.Version = aud.Version
		//检查是否添加过了
		user_app_info := &models.UserAppInfo{}
		code, _ := user_app_info.FindByAttribute(&uaf, 1)

		if code != 404 {
			err_info = app_error.ErrUserAppInfoExist
			break
		}
		createDate := strconv.FormatInt(time.Now().Unix(), 10)
		user_app := models.NewUserApp(&uaf, createDate)
		if _, err := user_app.Insert(); err != nil {
			err_info = app_error.ErrDatabase
			break
		}
	}
	//检查是否有错
	if err_info.Code != 0 {
		u.Data["json"] = err_info
		u.ServeJson()
		return
	}

	u.Data["json"] = app_error.SuccessData
	u.ServeJson()
}

// @Title 用户更新App接口
// @Description 用户更新App接口
// @Param	body		body 	models.UserAppForm	true	"用户更新app数据"
// @Success 200 {object} common.app_error.CodeInfo
// @Failure 403 body is empty
// @router /app_update [post]
func (u *UserAppController) AppUpdate() {
	user_app_form := models.UserAppFormSlice{}
	json.Unmarshal(u.Ctx.Input.RequestBody, &user_app_form)
	//*
	ids := map[int64]string{}
	//uafs := []models.UserAppForm{}
	err_info := app_error.CodeInfo{}
	// 检查是否重复
	for _, uaf := range user_app_form.UserAppForms {
		//uaf.(UserAppForm)
		//beego.Debug(uaf)
		//*
		if ids[uaf.AppId] == "1" {
			err_info = app_error.ErrDataDuplication
			//return
			break
		}
		//beego.Debug(uaf.AppId)

		ids[uaf.AppId] = "1"
		//获取应用的当前版本信息

		aud := app_upload.GetAppInfoById(uaf.AppId)
		//	beego.Debug("每次都有信息...", aud)
		if aud == nil {
			err_info = app_error.ErrUserAppInfoNoExist
			//return
			break
		}
		valid := validation.Validation{}
		ok, err := valid.Valid(&uaf)
		//如果存在错误
		if err != nil {
			err_info = app_error.ErrInputData
			break
		}
		if !ok {
			err_info = app_error.ErrDupUser
			break
		}
		uaf.Version = aud.Version
		//检查是否添加过了
		user_app_info := &models.UserAppInfo{}
		code, _ := user_app_info.FindByAttribute(&uaf, 1)

		if code == 404 {
			err_info = app_error.ErrUserAppInfoNoExist
			break
		}
		code, _ = user_app_info.UpdateVersion(&uaf)

		if code == -1 {
			err_info = app_error.ErrDatabase
			break
		}
	}
	//检查是否有错
	if err_info.Code != 0 {
		u.Data["json"] = err_info
		u.ServeJson()
		return
	}

	// update redis cache

	u.Data["json"] = app_error.SuccessData
	u.ServeJson()
}

// @Title 用户删除App接口
// @Description 用户删除App接口
// @Param	body		body 	models.UserAppForm	true	"用户删除App接口"
// @Success 200 {object} common.app_error.CodeInfo
// @Failure 403 body is empty
// @router /app_del [post]
func (u *UserAppController) AppRemove() {
	user_app_form := models.UserAppFormSlice{}
	json.Unmarshal(u.Ctx.Input.RequestBody, &user_app_form)
	//*
	ids := map[int64]string{}
	//uafs := []models.UserAppForm{}
	err_info := app_error.CodeInfo{}
	// 检查是否重复
	for _, uaf := range user_app_form.UserAppForms {
		//uaf.(UserAppForm)
		//beego.Debug(uaf)
		//*
		if ids[uaf.AppId] == "1" {
			err_info = app_error.ErrDataDuplication
			//return
			break
		}
		//beego.Debug(uaf.AppId)

		ids[uaf.AppId] = "1"
		//获取应用的当前版本信息

		aud := app_upload.GetAppInfoById(uaf.AppId)
		//	beego.Debug("每次都有信息...", aud)
		if aud == nil {
			err_info = app_error.ErrUserAppInfoNoExist
			//return
			break
		}
		valid := validation.Validation{}
		ok, err := valid.Valid(&uaf)
		//如果存在错误
		if err != nil {
			err_info = app_error.ErrInputData
			break
		}
		if !ok {
			beego.Debug("ValidRegsiterForm errors:")
			for _, err := range valid.Errors {
				beego.Debug(err.Key, err.Message)
			}
			err_info = app_error.ErrDupUser
			break
		}
		uaf.Version = aud.Version
		//检查是否添加过了
		user_app_info := &models.UserAppInfo{}
		code, _ := user_app_info.FindByAttribute(&uaf, 1)

		if code == 404 {
			err_info = app_error.ErrUserAppInfoNoExist
			break
		}
		code, _ = user_app_info.Remove()
		if code == -1 {
			err_info = app_error.ErrDatabase
			break
		}
	}
	//检查是否有错
	if err_info.Code != 0 {
		u.Data["json"] = err_info
		u.ServeJson()
		return
	}
	/*
		user_app_form := models.UserAppForm{}
		json.Unmarshal(u.Ctx.Input.RequestBody, &user_app_form)

		beego.Debug("ParseRegsiterForm:", &user_app_form)
		valid := validation.Validation{}
		ok, err := valid.Valid(&user_app_form)
		//如果存在错误
		if err != nil {
			beego.Debug("ValidRegsiterForm:", err)
			u.Data["json"] = app_error.ErrInputData
			u.ServeJson()
			return
		}
		if !ok {
			beego.Debug("ValidRegsiterForm errors:")
			for _, err := range valid.Errors {
				beego.Debug(err.Key, err.Message)
			}
			u.Data["json"] = app_error.ErrDupUser
			u.ServeJson()
			return
		}
		//检查数据是否存在
		user_app_info := &models.UserAppInfo{}
		code, _ := user_app_info.FindByAttribute(&user_app_form, 1)
		if code == 404 {
			beego.Debug("这里有问题:", user_app_info)
			u.Data["json"] = app_error.ErrUserAppInfoNoExist
			u.ServeJson()
			return
		}

		code, _ = user_app_info.Remove()

		if code == -1 {
			beego.Debug("InsertUserApp:", err)
			u.Data["json"] = app_error.ErrDatabase
			u.ServeJson()
			return
		}
	*/
	// update redis cache

	u.Data["json"] = app_error.SuccessData
	u.ServeJson()
}
