package logger

import (
    "fmt"
    "sync"
    "time"
    "runtime"
    "strings"
    "goweb/conf"
    "goweb/lib/tool"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "github.com/natefinch/lumberjack"
)

type tOption struct {
    MaxSize                     int
    LogPath                     string
    FileExt                     string
    LogType                     string
    TimeKey                     string
    NameTime                    string
    MessageKey                  string
    LogTimeZone                 *time.Location
    DefaultWriter               *lumberjack.Logger
}

var vOption tOption
var vInitLock sync.Mutex
var vLogWriterMap = make(map[string]map[string]*lumberjack.Logger)

func init() {
    initOption()
    initWriter()
    go initTimer()
}

func initTimer() {
    method := "App-Bar-Model-initTimer"
    ticker := time.NewTicker(time.Second * 1800)
    for {
        t := <-ticker.C
        s := t.Format("2006-01-02 15:04:05.999")
        initOption()
        initWriter()
        traceid := tool.GetTraceId()
        getZapLogObjByLogType("app").Infow(traceid, "func", method, "ticker", s)
    }
}

func initOption() {
    vInitLock.Lock()
    defer vInitLock.Unlock()

    fileext := ".log"
    logpath := conf.AppConfig.GetString("log.path")
    timezone, _ := time.LoadLocation("Asia/Shanghai")

    vOption = tOption{
        MaxSize: 1000000,
        LogPath: logpath,
        FileExt: fileext,
        LogType: "app|mysql|debug|error|request|response",
        TimeKey: "logtime",
        NameTime: "D",
        MessageKey: "traceid",
        LogTimeZone: timezone,
        DefaultWriter: &lumberjack.Logger{Filename: logpath + "/app/default" + fileext},
    }
}

func initWriter() {
    logtypes := strings.Split(vOption.LogType, "|")
    timenow := time.Now().In(vOption.LogTimeZone)
    maxsize := vOption.MaxSize
    names := []string{}

    if vOption.NameTime == "Y" {
        tset := []time.Time{
            timenow,
            timenow.AddDate(1, 0, 0),
        }
        for _, t := range tset {
            names = append(names, t.Format("2006"))
        }
    } else if vOption.NameTime == "M" {
        tset := []time.Time{
            timenow,
            timenow.AddDate(0, 1, 0),
        }
        for _, t := range tset {
            names = append(names, t.Format("200601"))
        }
    } else if vOption.NameTime == "D" {
        tset := []time.Time{
            timenow,
            timenow.AddDate(0, 0, 1),
        }
        for _, t := range tset {
            names = append(names, t.Format("20060102"))
        }
    } else if vOption.NameTime == "H" {
        tset := []time.Time{
            timenow,
            timenow.Add(time.Hour * 1),
        }
        for _, t := range tset {
            names = append(names, t.Format("2006010215"))
        }
    }

    vInitLock.Lock()
    defer vInitLock.Unlock()

    namem := make(map[string]bool)
    for _, name := range names { namem[name] = true }
    for _, logtype := range logtypes {
        wm := vLogWriterMap[logtype]
        for wname, writer := range wm { if !namem[wname] { writer.Close() } }
    }
    for _, name := range names {
        for _, logtype := range logtypes {
            filename := vOption.LogPath + "/" + logtype + "/" + name + vOption.FileExt
            if _, ok := vLogWriterMap[logtype]; !ok {
                vLogWriterMap[logtype] = make(map[string]*lumberjack.Logger)
            }
            if _, ok := vLogWriterMap[logtype][name]; !ok {
                vLogWriterMap[logtype][name] = &lumberjack.Logger{
                    LocalTime: true, MaxSize: maxsize, Compress: false, Filename: filename,
                }
            }
        }
    }
}

func getFileLine() (s string) {
    _, f, l, _ := runtime.Caller(3)

    return fmt.Sprintf("%s:%d", f, l)
}

func getZapLogObjByLogType(logtype string) (*zap.SugaredLogger) {
    file := vOption.DefaultWriter
    timenow := time.Now().In(vOption.LogTimeZone)

    w, ok := vLogWriterMap[logtype]
    if ok {
        var n string
        if vOption.NameTime == "Y" {
            n = timenow.Format("2006")
        } else if vOption.NameTime == "M" {
            n = timenow.Format("200601")
        } else if vOption.NameTime == "D" {
            n = timenow.Format("20060102")
        } else if vOption.NameTime == "H" {
            n = timenow.Format("2006010215")
        }
        if file, ok = w[n]; !ok { file = vOption.DefaultWriter }
    }

    logencoder := zapcore.NewJSONEncoder(
        zapcore.EncoderConfig{
            TimeKey: vOption.TimeKey,
            MessageKey: vOption.MessageKey,
            EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
                timezone := vOption.LogTimeZone
                enc.AppendString(t.In(timezone).Format("2006-01-02 15:04:05.999"))
            },
        },
    )

    return zap.New(zapcore.NewCore(logencoder, zapcore.AddSync(file), zapcore.InfoLevel)).Sugar()
}
