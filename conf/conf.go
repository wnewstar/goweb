package conf

import (
    "os"
    "fmt"
    "time"
    "github.com/spf13/viper"
    "github.com/fsnotify/fsnotify"
)

var AppConfig *viper.Viper
var TestConfig *viper.Viper

var vAppPath = "/code/go/goweb"
var vTimeFormat = "2006-01-02 15:04:05"
var vCnofNameMap = map[string]string{"deve": "deve", "test": "test", "prod": "prod", "prep": "prep"}

func init() {
    initApp()
    initTest()
}

/**
 * @desc    处理应用配置
 * @name    initApp
 * @date    2020-03-06
 * @author  wnewstar
 */
func initApp() {
    var ok bool
    var confname string
    var confpath = vAppPath + "/conf/file/app"
    if confname, ok = vCnofNameMap[os.Getenv("THS_TIER")]; !ok {
        confname = "deve"
    }

    AppConfig = viper.New()
    AppConfig.AddConfigPath(confpath)
    AppConfig.SetConfigName(confname)
    AppConfig.SetConfigType("json")
    err := AppConfig.ReadInConfig()
    if err != nil {
        fmt.Printf("[%s][error][%s]\n", time.Now().Format(vTimeFormat), err.Error())
    } else {
        AppConfig.WatchConfig()
        AppConfig.OnConfigChange(func(e fsnotify.Event) {
            fmt.Printf("[%s][config-change][%s]\n", time.Now().Format(vTimeFormat), e.Name)
        })
    }
}

/**
 * @desc    处理测试配置
 * @name    initTest
 * @date    2020-03-06
 * @author  wnewstar
 */
func initTest() {
    var ok bool
    var confname string
    var confpath = vAppPath + "/conf/file/test"
    if confname, ok = vCnofNameMap[os.Getenv("THS_TIER")]; !ok {
        confname = "deve"
    }

    TestConfig = viper.New()
    TestConfig.AddConfigPath(confpath)
    TestConfig.SetConfigName(confname)
    TestConfig.SetConfigType("json")
    err := TestConfig.ReadInConfig()
    if err != nil {
        fmt.Printf("[%s][error][%s]\n", time.Now().Format(vTimeFormat), err.Error())
    } else {
        TestConfig.WatchConfig()
        TestConfig.OnConfigChange(func(e fsnotify.Event) {
            fmt.Printf("[%s][config-change][%s]\n", time.Now().Format(vTimeFormat), e.Name)
        })
    }
}
