package model

import (
    "os"
    "fmt"
    "time"
    "goweb/conf"
    "goweb/lib/tool"
    "goweb/lib/logger"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

var VDbQuery *gorm.DB
var VDbWrite *gorm.DB

var VTimezone *time.Location

func init() {
    initDbMysql()
    VTimezone, _ = time.LoadLocation("Asia/Shanghai")
}

func initDbMysql() {
    var err error

    c := conf.AppConfig
    envt := c.GetString("env")
    sqllog := &logger.GormLogger{}
    method := "App-Bar-Model-initDbMysql"

    VDbQuery, err = gorm.Open("mysql", c.GetString("db.mysql.test.query.dsn"))
    if err != nil {
        logger.LogApp(tool.GetTraceId(), "func", method, "errmsg", err.Error())
        fmt.Println(err.Error())
        os.Exit(-1)
    } else {
        logger.LogApp(tool.GetTraceId(), "func", method, "database", "VDbQuery")
    }

    if envt == "deve" || envt == "test" {
        VDbQuery.LogMode(true)
        VDbQuery.SetLogger(sqllog)
    }
    VDbQuery.SingularTable(true)
    VDbQuery.DB().SetMaxIdleConns(c.GetInt("db.mysql.test.query.maxidle"))
    VDbQuery.DB().SetMaxOpenConns(c.GetInt("db.mysql.test.query.maxopen"))
    VDbQuery.DB().SetConnMaxLifetime(time.Duration(c.GetInt("db.mysql.test.write.maxlifetime")) * time.Second)

    VDbWrite, err = gorm.Open("mysql", c.GetString("db.mysql.test.write.dsn"))
    if err != nil {
        logger.LogApp(tool.GetTraceId(), "func", method, "errmsg", err.Error())
        fmt.Println(err.Error())
        os.Exit(-1)
    } else {
        logger.LogApp(tool.GetTraceId(), "func", method, "database", "VDbWrite")
    }

    if envt == "deve" || envt == "test" {
        VDbWrite.LogMode(true)
        VDbWrite.SetLogger(sqllog)
    }
    VDbWrite.SingularTable(true)
    VDbWrite.DB().SetMaxIdleConns(c.GetInt("db.mysql.test.write.maxidle"))
    VDbWrite.DB().SetMaxOpenConns(c.GetInt("db.mysql.test.write.maxopen"))
    VDbWrite.DB().SetConnMaxLifetime(time.Duration(c.GetInt("db.mysql.test.write.maxlifetime")) * time.Second)
}
