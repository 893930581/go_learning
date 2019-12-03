package mylogger

import (
	"fmt"
	"time"
)

type Logger struct {
	level logLevel
}

func NewLogger(loglev string) (Logger) {
	if lv,err := ConvertStrLevel(loglev); err!=nil{
		panic(err)
	}else{
		return Logger{lv}
	}
}


func log(level logLevel,format string,a ...interface{}) {
	msg := fmt.Sprintf(format,a...)
	now := time.Now()
	funcName, fileName, lineNum := getInfo(3)
	eLevel := ConvertLevelStr(level)
	fmt.Printf("[%s] [%v] [%s:%s:%d] %s\n",now.Format("2019-12-2 15:04:05"),eLevel,funcName, fileName, lineNum,msg)
}

func (l Logger) Debug(format string,a ...interface{}){
	if l.level <= DEBUG {
		log(DEBUG,format,a...)
	}
}

func (l Logger) Trace(format string,a ...interface{}){
	if l.level <= TRACE {
		log(TRACE,format,a...)
	}
}

func (l Logger) Info(format string,a ...interface{}){
	if l.level <= INFO {
		log(INFO,format,a...)
	}
}

func (l Logger) Warning(format string,a ...interface{}){
	if l.level <= WARNING {
		log(WARNING,format,a...)
	}
}

func (l Logger) Error(format string,a ...interface{}){
	if l.level <= ERROR {
		log(ERROR,format,a...)
	}
}

func (l Logger) Fatal(format string,a ...interface{}){
	if l.level <= FATAL {
		log(FATAL,format,a...)
	}
}
