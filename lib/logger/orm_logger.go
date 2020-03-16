package logger

import (
    "fmt"
    "goweb/lib/tool"
)

type GormLogger struct {}

func (l *GormLogger) Print(args ...interface{}) {
    LogMysql(too.GetTraceId(), l.formatGormInfo(args)...)
}

func (l *GormLogger) formatGormInfo(info []interface{}) ([]interface{}) {
    aset := []interface{}{}

    if (info[0] == "sql") {
        keymap := map[int]string {
            0: "infotype",
            1: "fileline",
            2: "timediff",
            3: "sql",
            4: "arg",
            5: "row",
        }
        for ak, av := range info {
        	nk, ok := keymap[ak]
            v := fmt.Sprint(av)
            if ok { aset = append(aset, nk, v) } else { aset = append(aset, fmt.Sprint(ak), v) }
        }
    }

    return aset
}
