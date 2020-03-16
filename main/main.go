package main

import (
    "fmt"
    "net/http"
    _ "net/http/pprof"
    "goweb/router"
    "github.com/gin-gonic/gin"
)

var hpgoweb string
var hppprof string
var appname string

func init() {
    hpgoweb = ":10000"
    hppprof = ":10001"
    appname = "bar-app"
}

func main() {
    gin.SetMode(gin.ReleaseMode)

    if len(hppprof) > 0 {
        go pprof()
    }

    fmt.Printf("%s-goweb start\n", appname)

    err := router.Route(gin.New()).Run(hpgoweb)
    if (err != nil) {
        fmt.Println(err)
        fmt.Printf("%s-goweb http service failed\n", appname)
    } else {
        fmt.Printf("%s-goweb http service success start %s\n", appname, hpgoweb)
    }
}

func pprof() {
    fmt.Printf("%s-pprof start\n", appname)

    err := http.ListenAndServe(hppprof, nil)
    if (err != nil) {
        fmt.Println(err)
        fmt.Printf("%s-pprof http service failed\n", appname)
    } else {
        fmt.Printf("%s-pprof http service success start %s\n", appname, hppprof)
    }
}
