package mylogger

import (
	"errors"
	"fmt"
	"io"
	"path"
	"runtime"
	"strings"
	"time"
)

type Logger interface {
	Debug(format string,a ...interface{})
	Trace(format string,a ...interface{})
	Info(format string,a ...interface{})
	Warning(format string,a ...interface{})
	Error(format string,a ...interface{})
	Fatal(format string,a ...interface{})
}

type logLevel uint16

const (
	UNKNOWN logLevel = iota
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

func ConvertStrLevel (levelStr string) (logLevel,error) {
	levelStr = strings.ToLower(levelStr)
	switch levelStr {
	case "debug":
		return DEBUG,nil
	case "trace":
		return TRACE,nil
	case "info":
		return INFO,nil
	case "warning":
		return WARNING,nil
	case "error":
		return ERROR,nil
	case "fatal":
		return FATAL,nil
	default:
		return UNKNOWN,errors.New("没有这种日志级别。")
	}
}

func ConvertLevelStr(level logLevel) (str string) {
	switch level {
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "DEBUG"
	}
}

func getInfo(skip int)(funcName,fileName string ,lineNum int){
	pc, file, lineNum, ok := runtime.Caller(skip)
	if !ok {
		fmt.Printf("Wrong Caller [%v]!!\n",skip)
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	return
}

func log(writer io.Writer,level logLevel,format string,a ...interface{}) {
	msg := fmt.Sprintf(format,a...)
	now := time.Now()
	funcName, fileName, lineNum := getInfo(3)
	eLevel := ConvertLevelStr(level)
	fmt.Fprintf(writer,"[%s] [%v] [%s:%s:%d] %s\n",now.Format("2019-12-2 15:04:05"),eLevel,funcName, fileName, lineNum,msg)
}