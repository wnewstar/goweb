package model

import (
    "time"
    "goweb/lib/logger"
)

type Hello struct {
    Id                          uint64                      `gorm:"column:tsn;primary_key;auto_increment"`
    Name                        string
}

func (*Hello) TableName() (string) {

    return "test"
}

func (*Hello) RecordLog(traceid string, mhello Hello) (uint64) {
    tss := time.Now().UnixNano() / 1e6

    VDbWrite.Create(&mhello)

    tse := time.Now().UnixNano() / 1e6
    method := "App-Bar-Model-Hello-RecordLog"
    logger.LogDebug(traceid, "class", "mysql", "func", method, "ts", tss, "td", tse - tss)

    return mhello.Id
}
