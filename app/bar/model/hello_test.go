package model

import (
    "testing"
    "goweb/lib/test"
)

func TestRecordLog(t *testing.T) {
    var mhello = Hello{Name: "test"}

    if mhello.RecordLog(test.GetTraceId(), mhello) == 0 { t.Error("[error][101]") }
}
