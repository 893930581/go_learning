package mylogger

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
)

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