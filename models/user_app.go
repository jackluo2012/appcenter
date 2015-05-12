package models

import (
	"appcenter/common/app_cache"
	. "appcenter/common/app_ckey"
	"appcenter/common/app_func"
	"appcenter/common/app_mongo"
	"appcenter/common/app_redis"
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"time"
)

const (
	UserAppTable = "user_apps"
)

type UserAppInfo struct {
	//ID         bson.ObjectId `bson:"_id,omitempty"`
	AppId      int64     `bson:"app_id"     json:"appid"`
	UserId     string    `bson:"user_id"      json:"user_id"`
	Udid       string    `bson:"udid" json:"udid"`
	Version    string    `bson:"version" json:"version"`
	CreateDate time.Time `bson:"create_date" json:"create_date"`
}

//查询数据
func GetAllUserApps(uasf *UserAppSearchForm) (code int, mlists []*UserAppInfo) {

	mConn := app_mongo.Conn()
	defer mConn.Close()
	c := mConn.DB("").C(UserAppTable)
	where := bson.M{}
	if (uasf.AppId != 0) && (uasf.Udid != "") {
		where = bson.M{"user_id": uasf.UserId, "app_id": uasf.AppId, "udid": uasf.Udid}
	} else if uasf.AppId != 0 {
		where = bson.M{"user_id": uasf.UserId, "app_id": uasf.AppId}
	} else if uasf.Udid != "" {
		where = bson.M{"user_id": uasf.UserId, "udid": uasf.Udid}
	} else {
		where = bson.M{"user_id": uasf.UserId}
	}
	err := c.Find(where).Skip((uasf.PageNo - 1) * uasf.PageSize).Limit(uasf.PageSize).All(&mlists)
	if err != nil {
		code = -1
	}

	return
}

/**
 *	根据 用户的udid,appid 获取相应的信息
 */
func GetUserAppsByUdid(uid string, udid string, appid int64) (uai *UserAppInfo) {

	if uai = GetUserAppsCacheByUdid(uid, udid, appid); uai == nil {
		user := UserAppSearchForm{UserId: uid, Udid: udid, AppId: appid}
		_, users := GetAllUserApps(&user)
		//udid_k := app_func.Md5([]byte(udid))
		if len(users) > 0 {
			for _, uai = range users {
				//beego.Debug(uai)
				SetUserAppsCacheByUdid(uid, udid, appid, uai)
			}

		}
	}

	return
}

/**
 *	存入redis中
 */
func GetUserAppsCacheByUdid(uid string, udid string, appid int64) (uai *UserAppInfo) {
	rConn := app_redis.Conn()
	defer rConn.Close()
	info := app_cache.CacheInfo{USERAPPLIST, []string{uid, app_func.Md5([]byte(udid)), strconv.FormatInt(appid, 10)}}
	key, _ := app_cache.GetKey(info)
	v, err := redis.Values(rConn.Do("HGETALL", key))
	if err != nil {
		panic(err)
	}

	if err := redis.ScanStruct(v, uai); err != nil {
		panic(err)
	}
	return
}

func SetUserAppsCacheByUdid(uid string, udid string, appid int64, uai *UserAppInfo) {
	rConn := app_redis.Conn()
	defer rConn.Close()
	//	udid_k := app_func.Md5([]byte(udid))
	info := app_cache.CacheInfo{USERAPPLIST, []string{uid, app_func.Md5([]byte(udid)), strconv.FormatInt(appid, 10)}}
	key, _ := app_cache.GetKey(info)
	beego.Debug(key)
	beego.Debug(uai)
	//*
	if _, err := rConn.Do("HMSET", redis.Args{}.Add(key).AddFlat(uai)...); err != nil {
		panic(err)
	}
	//*/
}

// 组织封装数据
func NewUserApp(apf *UserAppForm, t time.Time) *UserAppInfo {
	user_app_info := UserAppInfo{
		UserId:  apf.UserId,
		AppId:   apf.AppId,
		Udid:    apf.Udid,
		Version: apf.Version}
	user_app_info.CreateDate = t
	return &user_app_info
}

// 向mongodb 里面添加数据
func (uai *UserAppInfo) Insert() (code int, err error) {
	mConn := app_mongo.Conn()
	defer mConn.Close()
	c := mConn.DB("").C(UserAppTable)
	err = c.Insert(uai)
	code = 0
	if err != nil {
		code = -1
	}
	return
}

// 查询
func (uai *UserAppInfo) FindByAttribute(uaf *UserAppForm, v int) (code int, err error) {
	mConn := app_mongo.Conn()
	defer mConn.Close()
	c := mConn.DB("").C(UserAppTable)
	if v == 1 {
		err = c.Find(bson.M{"user_id": uaf.UserId, "app_id": uaf.AppId, "udid": uaf.Udid}).One(uai)
	} else {
		err = c.Find(bson.M{"user_id": uaf.UserId, "app_id": uaf.AppId, "udid": uaf.Udid, "version": uaf.Version}).One(uai)
	}
	if err != nil {
		if err == mgo.ErrNotFound {
			code = 404 //数据已不在
		} else {
			code = -1
		}
	} else {
		code = 0
	}

	return
}

// 修改  只改版本号
func (uai *UserAppInfo) UpdateVersion(uaf *UserAppForm) (code int, err error) {
	mConn := app_mongo.Conn()
	defer mConn.Close()
	c := mConn.DB("").C(UserAppTable)
	err = c.Update(bson.M{"user_id": uai.UserId, "app_id": uai.AppId, "udid": uai.Udid}, bson.M{"$set": bson.M{"version": uaf.Version}})

	if err != nil {
		if err == mgo.ErrNotFound {
			code = 404
		} else {
			code = -1
		}
	} else {
		code = 0
	}
	return
}

// 删除
func (uai *UserAppInfo) Remove() (code int, err error) {
	mConn := app_mongo.Conn()
	defer mConn.Close()
	c := mConn.DB("").C(UserAppTable)
	err = c.Remove(bson.M{"user_id": uai.UserId, "app_id": uai.AppId, "udid": uai.Udid})

	if err != nil {
		if err == mgo.ErrNotFound {
			code = 404
		} else {
			code = -1
		}
	} else {
		code = 0
	}
	return
}
