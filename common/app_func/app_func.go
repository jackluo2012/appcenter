package app_func

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"html/template"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func init() {
	beego.AddFuncMap("i18n", i18nHTML)
}

/**
 *	md5
 */
func Md5(buf []byte) string {
	hash := md5.New()
	hash.Write(buf)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func i18nHTML(lang, format string, args ...interface{}) template.HTML {
	return template.HTML(i18n.Tr(lang, format, args...))
}

/**
 * 获取目录
 */
func GetUploadPath(t string, appkey string) (url string) {
	//获取网站目录
	working_dir := beego.AppConfig.String("appuploadpath") + appkey
	url_path := beego.AppConfig.String("appurlpath") + appkey
	var filename string
	//relative_dir =
	switch t {
	case "icon":
		filename = getFileUrl(working_dir + "/icon*")
	case "pkg":
		filename = getFileUrl(working_dir + "/*.apk")
	case "zip":
		filename = getFileUrl(working_dir + "/*.zip")
	case "capsule":
		filename = getFileUrl(working_dir + "/*.capsule")
	case "showImg":
		filename = getFileUrl(working_dir + "/showImg.*")
	default:
		filename = ""

	}
	if t == "zip" {
		url = filename
	} else {
		url = url_path + strings.Replace(filename, working_dir, "", -1)
	}
	return
}

// 获取路径
func getFileUrl(file string) (path string) {
	files, err := filepath.Glob(file)
	if err != nil {
		return
	}
	if len(files) == 0 {
		return
	}
	return files[0]
}

/**
 *	get screens
 */

func GetUploadScreensPath(appkey string) (urls string) {
	//获取网站目录
	working_dir := beego.AppConfig.String("appuploadpath") + appkey
	url_path := beego.AppConfig.String("appurlpath") + appkey
	files, err := filepath.Glob(working_dir + "/screens/*")
	if err != nil {
		return
	}
	if len(files) == 0 {
		return
	}
	for _, filename := range files {
		filename = url_path + strings.Replace(filename, working_dir, "", -1)
		urls = urls + filename + "," //= append(urls, filename)
	}
	return
}

/**
 *	获取分类
 */
func CateTran(tp string) (ctp string) {
	tp = strings.Replace(tp, " ", "", -1)
	cateArr := map[string]string{"life": "生活",
		"education":  "教育",
		"business":   "商业",
		"entertainm": "娱乐",
		"fashion":    "时尚",
		"unkown":     "其它",
		"literature": "文学",
		"tour":       "旅游",
		"shoot":      "摄影",
		"music":      "音乐",
		"news":       "新闻",
		"medical":    "医疗",
		"social":     "社交",
		"catering":   "餐饮",
		"sport":      "运动",
		"health":     "健康",
		"tools":      "工具",
	}
	var err bool
	ctp, err = cateArr[tp]
	if !err {
		return tp
	}

	return
}

/**
 *	获取文件大小
 */
func GetFileSize(path string) (size string) {

	fileInfo, err := os.Stat(path)
	if err != nil {
		size = "0KB"
		return
	}
	size = strconv.FormatInt(int64(fileInfo.Size()/1024), 10) + "KB"
	return
}

/**
 *	获取下载地址
 */
func GetAppDownLoadUrl(appid int64) (url string) {

	url = beego.AppConfig.String("downloadurl") + strconv.FormatInt(appid, 10)
	return
}

/**
 * 验证密码
 */
func CheckSecurity(uid string, secretkey string) bool {

	sk := beego.AppConfig.String("securitykey")
	//beego.Debug("=========md5=========", Md5([]byte(uid+sk)))
	return Md5([]byte(uid+sk)) == secretkey
}

/**
 *	截取字符串
 */
func Substr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}
