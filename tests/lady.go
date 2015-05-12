package main

import (
	"log"

	"github.com/garyburd/redigo/redis"
	"github.com/garyburd/redigo/redisx"
)

type MyStruct struct {
	A int
	B string
}

type RequestLog struct {
	RequestType string //Get,Put,Post,Delete
	AccessKey   string
	SecretKey   string
	ReuqestURL  string //请求网址
	RemoteAddr  string
	CreateTime  string //创建日期
}

func main() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Fatal(err)
	}

	// v0 := &MyStruct{1, "hello"}
	v1 := &RequestLog{"GET", "a", "a", "/get/adsid/fasdfk", "19.2.23.2", "201403201528"}
	// _, err = c.Do("HMSET", redisx.AppendStruct([]interface{}{"key"}, v0)...)
	_, err = redisx.Do("HMSET", redisx.AppendStruct([]interface{}{"reqlog:201403201528.1"}, v1)...)
	if err != nil {
		log.Fatal(err)
	}

	reply, err := c.Do("HGETALL", "reqlog")
	if err != nil {
		log.Fatal(err)
	}

	v2 := &RequestLog{}

	err = redisx.ScanStruct(reply, v2)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("v2=%v", v2)
}
