package controller

import (
    "goweb/lib/logger"
    "github.com/gin-gonic/gin"
    ModelBar "goweb/app/bar/model"
)

type Hello struct {}

func (h *Hello) Hello(c *gin.Context) {
    var Mhello = ModelBar.Hello{}

    traceid := getTraceId(c)
    Mhello.Name = "测试数据"

    data := Mhello.RecordLog(traceid, Mhello)

    c.JSON(
        CODE_HTTP_SUCCESS,
        gin.H{"code": CODE_BASE_SUCCESS_A, "msg": "success", "traceid": traceid},
    )

    logger.LogResponse(traceid, "code", CODE_BASE_SUCCESS_A, "msg", "success", "data", data)
}
