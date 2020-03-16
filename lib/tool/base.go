package tool

import (
    "fmt"
    "time"
)

func GetTraceId() (string) {

    return "application" + time.Now().Format("20060102150405") + GetRandNumStr(7)
}
