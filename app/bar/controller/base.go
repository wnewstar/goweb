package controller

import (
    "io"
    "os"
    "time"
    "path"
    "errors"
    "goweb/conf"
    "github.com/gin-gonic/gin"
)

const (
    CODE_HTTP_SUCCESS = 200

    // 成功
    CODE_BASE_SUCCESS_A = 1
    CODE_BASE_FALIURE_A = 1100
    CODE_BASE_FALIURE_B = 1200
    CODE_BASE_FALIURE_C = 1300
    CODE_BASE_FALIURE_D = 1400
    CODE_BASE_FALIURE_E = 1500
)

func getTraceId(c *gin.Context) (string) {

    return c.Request.Header.Get("X-TraceId")
}

/**
 * @desc    从请求表单中保存邮件附件
 *          附件名和TRACEID确定唯一
 * @name    saveAttachment
 * @date    2020-01-09
 * @author  wnewstar
 * @param   c               *gin.Context                    请求
 * @param   data            string                          表单字段
 * @return  newname         string                          新的附件地址
 */
func saveAttachment(c *gin.Context, name, data string) (newname string, err error) {
    base := conf.AppConfig.GetString("attachment.path") + time.Now().Format("/20060102/")
    size := conf.AppConfig.GetInt64("attachment.size") * 1024 * 1024

    fext := path.Ext(name)
    file, err := c.FormFile(data)
    if err == nil {
        newname = base + getTraceId(c) + "-" + data + fext
        if file.Size > size {
            err = errors.New("附件文件大小超出限制-" + data)
        }
        if err == nil {
            err = os.MkdirAll(base, os.ModePerm)
        }
        if err == nil { err = c.SaveUploadedFile(file, newname) }
    } else {
        // 上传的不是文件，判断是不是字符串
        info := c.PostForm(data)
        newname = base + getTraceId(c) + "-" + data + fext
        if len(info) > 0 {
            err = nil
        }
        if int64(len(info)) > size {
            err = errors.New("附件文件大小超出限制-" + data)
        }
        if err == nil {
            err = os.MkdirAll(base, os.ModePerm)
        }
        if err == nil {
            var temp *os.File
            defer temp.Close()
            if temp, err = os.Create(newname); err == nil { _, err = io.WriteString(temp, info) }
        }
    }

    return
}
