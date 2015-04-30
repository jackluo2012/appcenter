package app_cache

import (
	"errors"
	"strconv"
)

const (
	Split = ":"
)

type CacheInfo struct {
	Key    string
	Subkey interface{}
}

func GetKey(ci CacheInfo) (string, err error) {
	key := ci.Key + Split
	//断言
	if ci.Subkey != nil {
		switch c := ci.Subkey.(type) {

		case int: //如果是数字
			key = key + Split + strconv.Itoa(c) + Split
		case []string:
			var tmp string
			for _, u := range c {
				// 将接口转换成string
				tmp += u + Split
			}
			key = key + tmp
		default:
			return nil, errors.New("subkey type error!!!")
		}
	}

	return key, nil
}
