package logger

func LogApp(msg string, kvs ...interface{}) {
    writeLogByLogType("app", msg, kvs...)
}

func LogMysql(msg string, kvs ...interface{}) {
    writeLogByLogType("mysql", msg, kvs...)
}

func LogDebug(msg string, kvs ...interface{}) {
    writeLogByLogType("debug", msg, kvs...)
}

func LogError(msg string, kvs ...interface{}) {
    writeLogByLogType("error", msg, kvs...)
}

func LogRequest(msg string, kvs ...interface{}) {
    writeLogByLogType("request", msg, kvs...)
}

func LogResponse(msg string, kvs ...interface{}) {
    writeLogByLogType("response", msg, kvs...)
}

func writeLogByLogType(logtype string, msg string, kvs ...interface{}) {
    kvs = append(kvs, "fileline", getFileLine())
    getZapLogObjByLogType(logtype).Infow(msg, append(kvs, "logtype", logtype)...)
}
