package router

import (
    "goweb/lib/plugin"
    "github.com/gin-gonic/gin"
    ControllerBar "goweb/app/bar/controller"
)

func ok(c *gin.Context) {
    c.JSON(200, gin.H{"code": 1, "msg": "ok"})
}

func Route(r *gin.Engine) (*gin.Engine) {
    r.Use(plugin.DoLogger())
    r.Use(plugin.SetTraceId())

    r.Any("/health", ok)

    controller := &ControllerBar.Hello{}
    r.POST("/bar/hello/hello", controller.Hello)

    return r
}
