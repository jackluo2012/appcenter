package app_error

type CodeInfo struct {
	Code int
	Info string
}

var (
	SuccessData           = CodeInfo{Code: 1111, Info: "success"}
	ErrInputData          = CodeInfo{Code: 1000, Info: "数据输入错误"}
	ErrDupUser            = CodeInfo{Code: 1001, Info: "数据格式错误"}
	ErrNoUser             = CodeInfo{Code: 1002, Info: "用户信息不存在"}
	ErrPass               = CodeInfo{Code: 1003, Info: "密码不正确"}
	ErrNoUserPass         = CodeInfo{Code: 1004, Info: "用户信息不存在或密码不正确"}
	ErrInvalidUser        = CodeInfo{Code: 1005, Info: "用户信息不正确"}
	ErrDatabase           = CodeInfo{Code: 1006, Info: "数据库操作错误"}
	ErrUserAppInfoExist   = CodeInfo{Code: 1007, Info: "用户App信息已存在"}
	ErrUserAppInfoNoExist = CodeInfo{Code: 1008, Info: "用户App信息不存在,更新失败"}
	ErrDataDuplication    = CodeInfo{Code: 1009, Info: "提交数据中存在重复的数据,添加失败"}
)
