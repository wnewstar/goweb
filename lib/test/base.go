package test

import (
    "fmt"
    "time"
    "math/rand"
    "goweb/conf"
)

// 单元测试获取配置信息
var AppConfig = conf.AppConfig
var TestConfig = conf.TestConfig

func GetTraceId() (string) {

    return "unittest" + time.Now().Format("20060102150405") + GetRandNumStr(10)
}
