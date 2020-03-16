package plugin

import (
    "os"
    "fmt"
    "time"
    "strconv"
    "crypto/md5"
    "encoding/hex"
    "goweb/lib/tool"
    "goweb/lib/logger"
    "github.com/gin-gonic/gin"
)

type gct = gin.Context
type ghf = gin.HandlerFunc

var host, _ = os.Hostname()

/**
 * @desc    恢复
 * @name    Recover
 * @date    2020-01-02
 * @author  wnewstar
 */
func Recover(c *gct) {
    if err := recover(); err != nil {
        c.String(500, "程序内部错误")

        r := c.Request

        kvs := []interface{}{}
        kvs = append(kvs, "host", host)
        kvs = append(kvs, "errmsg", fmt.Sprint(err))

        logger.LogError(r.Header.Get("X-TraceId"), kvs...)
    }
}

/**
 * @desc    记录日志
 * @name    DoLogger
 * @date    2020-01-02
 * @author  wnewstar
 */
func DoLogger() (ghf) {
    return func (c *gct) {
        defer Recover(c)

        tss := time.Now().UnixNano() / 1e6
        c.Next()
        tse := time.Now().UnixNano() / 1e6

        r := c.Request

        kvs := []interface{}{}
        kvs = append(kvs, "ts", tss)
        kvs = append(kvs, "td", tse - tss)
        kvs = append(kvs, "ip", c.ClientIP())
        kvs = append(kvs, "url", r.URL.Path)
        kvs = append(kvs, "host", host)
        kvs = append(kvs, "method", r.Method)
        kvs = append(kvs, "status", c.Writer.Status())

        logger.LogRequest(r.Header.Get("X-TraceId"), kvs...)
    }
}

/**
 * @desc    设置TRACEID
 * @name    SetTraceId
 * @date    2020-01-06
 * @author  wnewstar
 */
func SetTraceId() (ghf) {
    return func(c *gct) {
        if len(c.Request.Header.Get("X-TraceId")) == 0 {
            rnum := tool.GetRandNumStr(8)
            unix := time.Now().UnixNano()
            nano := strconv.FormatInt(unix, 10)
            data := md5.Sum([]byte(host + "-" + nano + "-" + rnum))

            // 主机 + 时间 + 随机数
            c.Request.Header.Set("X-TraceId", hex.EncodeToString(data[:]))
        }
    }
}
